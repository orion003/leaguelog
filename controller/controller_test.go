package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"leaguelog/database/mock"
	"leaguelog/logging"
	"leaguelog/model"
)

var server *httptest.Server
var controller *Controller

func init() {
	log := logging.NewLog15()

	controller = NewController(log)
	initializeRepos()

	server = httptest.NewServer(NewRouter(controller, ""))
}

func TestAddEmail(t *testing.T) {
	err := addEmail("test@test.com")
	if err != nil {
		t.Error(err)
	}
}

func TestDuplicateEmail(t *testing.T) {
	email := "test2@test.com"

	err := addEmail(email)
	if err != nil {
		t.Error(err)
	}

	err = addEmail(email)
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
	l1 := &model.League{
		Name:  "Test League 1",
		Sport: "Sport 1",
	}
	err := controller.leagueRepo.Create(l1)

	l2 := &model.League{
		Name:  "Test League 2",
		Sport: "Sport 2",
	}
	err = controller.leagueRepo.Create(l2)

	body, err := request("/api/leagues")
	if err != nil {
		t.Errorf("Unable to obtain body: %v", err)
	}

	leagues := make([]model.League, 2)
	if len(body) > 0 {
		err = json.Unmarshal(body, &leagues)
		if err != nil {
			fmt.Printf("Unable to unmarshal JSON: %v\n", err)
		}
	}

	if len(leagues) != 2 {
		t.Errorf("Incorrect number of leagues. Found %d, should be %d", len(leagues), 2)
	}
}

func TestGetStandings(t *testing.T) {
	l := &model.League{
		Name:  "Test League 1",
		Sport: "Sport 1",
	}
	err := controller.leagueRepo.Create(l)

	t1 := &model.Team{
		Name:   "Test Team 1",
		League: l,
	}
	err = controller.teamRepo.Create(t1)

	t2 := &model.Team{
		Name:   "Test Team 2",
		League: l,
	}
	err = controller.teamRepo.Create(t2)

	startDate := time.Date(2015, time.October, 6, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2016, time.April, 24, 0, 0, 0, 0, time.UTC)
	season := &model.Season{
		League:    l,
		Name:      "Test Season 1",
		StartDate: startDate,
		EndDate:   endDate,
	}
	err = controller.seasonRepo.Create(season)

	s1 := &model.Standing{
		Season: season,
		Team:   t1,
		Wins:   2,
		Losses: 1,
		Ties:   1,
	}
	err = controller.standingRepo.Create(s1)

	s2 := &model.Standing{
		Season: season,
		Team:   t1,
		Wins:   1,
		Losses: 2,
		Ties:   1,
	}
	err = controller.standingRepo.Create(s2)

	body, err := request(fmt.Sprintf("/api/league/%d/standings", l.Id))
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

	for _, standing := range standings {
		teamName := standing.Team.Name
		if teamName != t1.Name && teamName != t2.Name {
			t.Errorf("Team name does not match either team: %s", teamName)
		}
	}
}

func TestGetSchedule(t *testing.T) {
	l := &model.League{
		Name:  "Test League 1",
		Sport: "Sport 1",
	}
	err := controller.leagueRepo.Create(l)

	t1 := &model.Team{
		Name:   "Test Team 1",
		League: l,
	}
	err = controller.teamRepo.Create(t1)

	t2 := &model.Team{
		Name:   "Test Team 2",
		League: l,
	}
	err = controller.teamRepo.Create(t2)

	nextYear := time.Now().Year() + 1

	startDate := time.Date(nextYear, time.October, 6, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(nextYear+1, time.April, 24, 0, 0, 0, 0, time.UTC)
	season := &model.Season{
		League:    l,
		Name:      "Test Season 1",
		StartDate: startDate,
		EndDate:   endDate,
	}
	err = controller.seasonRepo.Create(season)

	startDate1 := time.Date(nextYear, time.October, 6, 20, 30, 0, 0, time.UTC)
	g1 := &model.Game{
		Season:    season,
		StartTime: startDate1,
		HomeTeam:  t1,
		AwayTeam:  t2,
		HomeScore: 0,
		AwayScore: 0,
	}
	err = controller.gameRepo.Create(g1)

	startDate2 := time.Date(nextYear, time.October, 8, 19, 15, 0, 0, time.UTC)
	g2 := &model.Game{
		Season:    season,
		StartTime: startDate2,
		HomeTeam:  t2,
		AwayTeam:  t1,
		HomeScore: 0,
		AwayScore: 0,
	}
	err = controller.gameRepo.Create(g2)

	body, err := request(fmt.Sprintf("/api/league/%d/schedule", l.Id))
	if err != nil {
		t.Errorf("Unable to obtain body: %v", err)
	}

	games := []model.Game{}
	if len(body) > 0 {
		err = json.Unmarshal(body, &games)
		if err != nil {
			t.Errorf("Unable to unmarshal JSON: %v\n", err)
		}
	}

	expectedCount := 2
	if len(games) != expectedCount {
		t.Errorf("Incorrect number of games. Found %d, should be %d", len(games), expectedCount)
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

func initializeRepos() {
	leagueRepo := mock.NewMockLeagueRepository()
	seasonRepo := mock.NewMockSeasonRepository()
	teamRepo := mock.NewMockTeamRepository()
	standingRepo := mock.NewMockStandingRepository()
	gameRepo := mock.NewMockGameRepository()
	userRepo := mock.NewMockUserRepository()

	controller.SetLeagueRepository(leagueRepo)
	controller.SetSeasonRepository(seasonRepo)
	controller.SetTeamRepository(teamRepo)
	controller.SetStandingRepository(standingRepo)
	controller.SetGameRepository(gameRepo)
	controller.SetUserRepository(userRepo)
}
