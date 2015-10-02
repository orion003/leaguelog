package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"recleague/database/mock"
	"recleague/logging"
	"recleague/model"
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
	u := fmt.Sprintf("%s/api/users", server.URL)

	v := url.Values{}
	v.Set("email", email)

	res, err := http.DefaultClient.PostForm(u, v)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	controller.log.Info(fmt.Sprintf("HTTP Status: %d", res.StatusCode))

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

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
        Name: "Test League 1",
        Sport: "Sport 1",
    }
    err := controller.leagueRepo.Create(l1)
    
    l2 := &model.League{
        Name: "Test League 2",
        Sport: "Sport 2",
    }
    err = controller.leagueRepo.Create(l2)
    
    u := fmt.Sprintf("%s/api/leagues", server.URL)

	res, err := http.DefaultClient.Get(u)
	if err != nil {
		fmt.Printf("Unable to access %s: %v\n", err);
	}

	defer res.Body.Close()
	
	controller.log.Info(fmt.Sprintf("HTTP Status: %d", res.StatusCode))

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read body: %v\n", u, err);
	}
	
    leagues := make([]model.League, 2)
	if len(body) > 0 {
		err = json.Unmarshal(body, &leagues)
		if err != nil {
			fmt.Printf("Unable to unmarshal JSON: %v\n", err);
		}
	}
	
	if(len(leagues) != 2) {
	    t.Errorf("Incorrect number of leagues. Found %d, should be %d", len(leagues), 2)   
	}
}

func initializeRepos() {
	leagueRepo := mock.NewMockLeagueRepository()
	seasonRepo := mock.NewMockSeasonRepository()
	teamRepo := mock.NewMockTeamRepository()
	gameRepo := mock.NewMockGameRepository()
	userRepo := mock.NewMockUserRepository()

	controller.SetLeagueRepository(leagueRepo)
	controller.SetSeasonRepository(seasonRepo)
	controller.SetTeamRepository(teamRepo)
	controller.SetGameRepository(gameRepo)
	controller.SetUserRepository(userRepo)
}
