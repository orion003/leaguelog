package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"leaguelog/auth/service"
	"leaguelog/database/postgres"
	"leaguelog/logging"
	"leaguelog/model"
)

var log = logging.NewLog15()

var server *httptest.Server
var controller *Controller

var c config

type config struct {
	Database database `json:"database"`
}

type database struct {
	URL  string `json:"url"`
	Seed string `json:"seed"`
	Test string `json:"test"`
}

func TestMain(m *testing.M) {
	controller = NewController(log)
	initialize()

	server = httptest.NewServer(NewRouter(controller, ""))

	r := m.Run()

	server.Close()
	controller.repo.Close()
	os.Exit(r)
}

func TestGetLeagues(t *testing.T) {
	body, err := request("/api/leagues")
	if err != nil {
		t.Errorf("Unable to obtain body: %v", err)
	}

	leagues := make([]model.League, 1)
	if len(body) > 0 {
		err = json.Unmarshal(body, &leagues)
		if err != nil {
			fmt.Printf("Unable to unmarshal JSON: %v\n", err)
		}
	}

	if len(leagues) != 1 {
		t.Errorf("Incorrect number of leagues. Found %d, should be %d", len(leagues), 1)
	}
}

func TestGetStandings(t *testing.T) {
	leagueID := 1

	body, err := request(fmt.Sprintf("/api/league/%d/standings", leagueID))
	if err != nil {
		t.Errorf("Unable to obtain body: %v", err)
	}

	standings := []model.Standing{}
	if len(body) > 0 {
		err = json.Unmarshal(body, &standings)
		if err != nil {
			t.Errorf("Unable to unmarshal JSON: %v\n", err)
		}
	}

	expectedCount := 2
	if len(standings) != expectedCount {
		t.Errorf("Incorrect number of standings. Found %d, should be %d", len(standings), expectedCount)
	}
}

func TestGetSchedule(t *testing.T) {
	seasonID := 1
	season, err := controller.repo.FindSeasonByID(seasonID)
	if err != nil {
		t.Errorf("Unable to find season: %v", err)
	}

	homeTeamID := 1
	homeTeam, err := controller.repo.FindTeamByID(homeTeamID)
	if err != nil {
		t.Errorf("Unable to find home team: %v", err)
	}

	awayTeamID := 1
	awayTeam, err := controller.repo.FindTeamByID(awayTeamID)
	if err != nil {
		t.Errorf("Unable to find away team: %v", err)
	}

	year, month, day := time.Now().AddDate(0, 0, 1).Date()
	startTime := time.Date(year, month, day, 20, 30, 0, 0, time.UTC)
	game := &model.Game{
		Season:    season,
		StartTime: startTime,
		HomeTeam:  homeTeam,
		AwayTeam:  awayTeam,
		HomeScore: 0,
		AwayScore: 0,
	}
	err = controller.repo.CreateGame(game)

	leagueID := 1
	body, err := request(fmt.Sprintf("/api/league/%d/schedule", leagueID))
	if err != nil {
		t.Errorf("Unable to obtain body: %v", err)
	}

	schedules := []Schedule{}
	if len(body) > 0 {
		err = json.Unmarshal(body, &schedules)
		if err != nil {
			t.Errorf("Unable to unmarshal JSON: %v\n", err)
		}
	}

	expectedCount := 1
	if len(schedules) != expectedCount {
		t.Errorf("Incorrect number of games. Found %d, should be %d", len(schedules), expectedCount)
	}

	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	if schedules[0].StartDate != startDate {
		t.Errorf("Incorrect date for schedule - Expected: %v, Received: %v", schedules[0].StartDate, startDate)
	}
}

func request(rURL string) ([]byte, error) {
	u := fmt.Sprintf("%s%s", server.URL, rURL)
	controller.log.Info(fmt.Sprintf("Test Request: %s", u))

	res, err := http.DefaultClient.Get(u)
	if err != nil {
		fmt.Printf("Unable to access %s: %v\n", u, err)
		return []byte{}, err
	}

	defer res.Body.Close()

	controller.log.Info(fmt.Sprintf("HTTP Status: %d", res.StatusCode))

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read body: %v\n", err)
		return []byte{}, err
	}

	return body, nil
}

func post(rURL string, data string) ([]byte, error) {
	u := fmt.Sprintf("%s%s", server.URL, rURL)

	res, err := http.DefaultClient.Post(u, "application/json", strings.NewReader(data))
	if err != nil {
		return []byte{}, err
	}

	controller.log.Info(fmt.Sprintf("HTTP Status: %d", res.StatusCode))

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func initialize() error {
	file, err := ioutil.ReadFile("./_config.json")
	if err != nil {
		log.Error(fmt.Sprintf("Unable to open file: %s", "./_config.json"))
		log.Error(fmt.Sprintf("Unable to initialize tests: %v", err))
		os.Exit(1)
	}

	json.Unmarshal(file, &c)

	err = initializeDatabase()
	if err != nil {
		log.Error(fmt.Sprintf("Unable to initialize tests: %v", err))
		os.Exit(1)
	}

	repo := initializeRepo()
	controller.SetRepository(repo)

	token := initializeToken()
	controller.SetTokenService(token)

	return nil
}

func initializeDatabase() error {
	db, err := sql.Open("postgres", c.Database.URL+c.Database.Seed)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Info(fmt.Sprintf("Dropping database: %s", c.Database.Test))
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", c.Database.Test))
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Creating database: %s from template %s", c.Database.Test, c.Database.Seed))
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s TEMPLATE %s", c.Database.Test, c.Database.Seed))
	if err != nil {
		return err
	}

	return nil
}

func initializeRepo() model.Repository {
	url := c.Database.URL + c.Database.Test

	manager, err := postgres.NewPgManager(url)
	if err != nil {
		log.Error("Unable to initialize PgManager")
		os.Exit(1)
	}

	return postgres.NewPgRepository(manager)
}

func initializeToken() service.TokenService {
	var hmac = []byte("579760E50509F2F28324421C7509741F5BF03B9158161076B3C6B39FB028D9E2C251490A3F8BD1F59728259A673668CAFEB49C9E9499F8386B147D7260B6937A")

	return service.InitializeJwt(hmac)
}
