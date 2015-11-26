package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"leaguelog/database/postgres"
	"leaguelog/logging"
	"leaguelog/model"
)

var log logging.Logger = logging.NewLog15()

var server *httptest.Server
var controller *Controller

var c config

type config struct {
	Database database `json:"database"`
}

type database struct {
	Url  string `json:"url"`
	Seed string `json:"seed"`
	Test string `json:"test"`
}

func TestMain(m *testing.M) {
	controller = NewController(log)
	initialize(controller)

	server = httptest.NewServer(NewRouter(controller, ""))

	r := m.Run()

	os.Exit(r)
}

func TestAuthController(t *testing.T) {
    testUserRegister(t *testing.T)   
}

func TestAddEmail(t *testing.T) {
	email := "test2@leaguelog.ca"
	err := addEmail(email)
	if err != nil {
		t.Errorf("Unable to add email: %s - %v", email, err)
	}
}

func TestDuplicateEmail(t *testing.T) {
	email := "test@leaguelog.ca"
	err := addEmail(email)
	if err == nil {
		t.Errorf("Duplicate email not allowed: %s", email)
	}

	if err.Error() != model.UserDuplicateEmail.Error() {
		t.Errorf("Should be duplicate email error: %s", err)
	}
}

func TestInvalidEmail(t *testing.T) {
	email := "test_invalid"

	err := addEmail(email)
	if err == nil {
		t.Errorf("Invalid email not allowed: %s", email)
	}

	if err.Error() != model.UserInvalidEmail.Error() {
		t.Errorf("Should be invalid email error: %s", err)
	}
}

func addEmail(email string) error {
	data := fmt.Sprintf("{\"email\": \"%s\"}", email)

	body, err := post("/api/users", data)

	if len(body) > 0 {
		var r map[string]interface{}
		err = json.Unmarshal(body, &r)
		if err != nil {
			return err
		}

		if e, ok := r["error"]; ok {
			if s, ok := e.(string); ok {
				return errors.New(s)
			}

			return errors.New("String value not found for json error.")
		}
	}

	return nil
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
	leagueId := 1

	body, err := request(fmt.Sprintf("/api/league/%d/standings", leagueId))
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
	seasonId := 1
	season, err := controller.repo.FindSeasonById(seasonId)
	if err != nil {
		t.Errorf("Unable to find season: %v", err)
	}

	homeTeamId := 1
	homeTeam, err := controller.repo.FindTeamById(homeTeamId)
	if err != nil {
		t.Errorf("Unable to find home team: %v", err)
	}

	awayTeamId := 1
	awayTeam, err := controller.repo.FindTeamById(awayTeamId)
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

	leagueId := 1
	body, err := request(fmt.Sprintf("/api/league/%d/schedule", leagueId))
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

func request(rUrl string) ([]byte, error) {
	u := fmt.Sprintf("%s%s", server.URL, rUrl)
	controller.log.Info(fmt.Sprintf("Test Request: %s", u))

	res, err := http.DefaultClient.Get(u)
	if err != nil {
		fmt.Printf("Unable to access %s: %v\n", err)
		return []byte{}, err
	}

	defer res.Body.Close()

	controller.log.Info(fmt.Sprintf("HTTP Status: %d", res.StatusCode))

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read body: %v\n", u, err)
		return []byte{}, err
	}

	return body, nil
}

func post(rUrl string, data string) ([]byte, error) {
	u := fmt.Sprintf("%s%s", server.URL, rUrl)

	res, err := http.DefaultClient.Post(u, "application/json", strings.NewReader(data))
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

	controller.log.Info(fmt.Sprintf("HTTP Status: %d", res.StatusCode))

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func initialize(controller *Controller) error {
	file, err := ioutil.ReadFile("./_config.json")
	if err != nil {
		return errors.New("Unable to open file: _config.json")
	}

	json.Unmarshal(file, &c)

	err = initializeTables()
	if err != nil {
		return err
	}

	repo := initializeRepo()
	controller.SetRepository(repo)

	return nil
}

func initializeTables() error {
	db, err := sql.Open("postgres", c.Database.Url+c.Database.Seed)
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
	url := c.Database.Url + c.Database.Test

	manager, err := postgres.NewPgManager(url)
	if err != nil {
		log.Error("Unable to initialize PgManager")
		os.Exit(1)
	}

	return postgres.NewPgRepository(manager)
}
