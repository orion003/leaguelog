package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"leaguelog/model"
)

type testLogin struct {
	email    string
	password string
	token    string
	err      error
}

var registrations = []testLogin{
	testLogin{
		email:    "test_register@leaguelog.ca",
		password: "test_password",
		err:      nil,
	},
	testLogin{
		email:    "test@leaguelog.ca",
		password: "duplicate_password",
		err:      model.UserDuplicateEmail,
	},
	testLogin{
		email:    "test_invalid",
		password: "invalid_password",
		err:      model.UserInvalidEmail,
	},
	testLogin{
		email:    "test_missing_password@leaguelog.ca",
		password: "",
		err:      model.UserInvalidPassword,
	},
}

var logins = []testLogin{
	testLogin{
		email:    "test@leaguelog.ca",
		password: "password",
		err:      nil,
	},
	testLogin{
		email:    "test@leaguelog.ca",
		password: "incorrect_password",
		err:      model.UserIncorrectPassword,
	},
	testLogin{
		email:    "test_unknown@leaguelog.ca",
		password: "unknown_password",
		err:      model.UserUnknownEmail,
	},
}

func TestUserRegister(t *testing.T) {
	for _, reg := range registrations {
		data := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, reg.email, reg.password)

		err := sendRegisterRequest(data)

		if err == nil && reg.err != nil {
			t.Errorf("Error should have been received: %v", reg.err)
		}

		if err != nil && reg.err == nil {
			t.Errorf("Error should not have been received: %v", err)
		}

		if err != nil && reg.err != nil && err.Error() != reg.err.Error() {
			t.Errorf("Incorrect error received. Expected: %v, Received: %v", reg.err, err)
		}
	}
}

func TestUserLogin(t *testing.T) {
	for _, l := range logins {
		data := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, l.email, l.password)

		err := sendLoginRequest(data)

		if err == nil && l.err != nil {
			t.Errorf("Error should have been received: %v", l.err)
		}

		if err != nil && l.err == nil {
			t.Errorf("Error should not have been received: %v", err)
		}

		if err != nil && l.err != nil && err.Error() != l.err.Error() {
			t.Errorf("Incorrect error received. Expected: %v, Received: %v", l.err, err)
		}
	}
}

func sendRegisterRequest(data string) error {
	return sendAuthRequest("/api/user/register", data)
}

func sendLoginRequest(data string) error {
	return sendAuthRequest("/api/user/login", data)
}

func sendAuthRequest(url string, data string) error {
	body, err := post(url, data)
	if err != nil {
		return fmt.Errorf("Unsuccessful post: %v\n", err)
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
	} else {
		fmt.Println("Empty body found when registering user.")
	}

	return nil
}
