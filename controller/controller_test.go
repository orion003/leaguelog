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

	server = httptest.NewServer(NewRouter(controller))
}

func TestAddEmail2(t *testing.T) {
	err := addEmail("test3@test.com")
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
	u := fmt.Sprintf("%s/users", server.URL)

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
