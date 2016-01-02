package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
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
		token:    "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NTAyODY4MTgsImlhdCI6MTQ1MDAyNzYxOCwiaXNzIjoibGVhZ3VlbG9nLmNhIiwidWlkIjoidGVzdEBsZWFndWVsb2cuY2EifQ.",
		err:      nil,
	},
	testLogin{
		email:    "test@leaguelog.ca",
		password: "duplicate_password",
		token:    "",
		err:      model.UserDuplicateEmail,
	},
	testLogin{
		email:    "test_invalid",
		password: "invalid_password",
		token:    "",
		err:      model.UserInvalidEmail,
	},
	testLogin{
		email:    "test_missing_password@leaguelog.ca",
		password: "",
		token:    "",
		err:      model.UserInvalidPassword,
	},
}

var logins = []testLogin{
	testLogin{
		email:    "test@leaguelog.ca",
		password: "password",
		token:    "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NTAyODY4MTgsImlhdCI6MTQ1MDAyNzYxOCwiaXNzIjoibGVhZ3VlbG9nLmNhIiwidWlkIjoidGVzdEBsZWFndWVsb2cuY2EifQ.",
		err:      nil,
	},
	testLogin{
		email:    "test@leaguelog.ca",
		password: "incorrect_password",
		token:    "",
		err:      model.UserIncorrectPassword,
	},
	testLogin{
		email:    "test_unknown@leaguelog.ca",
		password: "unknown_password",
		token:    "",
		err:      model.UserUnknownEmail,
	},
}

func TestUserRegister(t *testing.T) {
	for _, reg := range registrations {
		testUserAction(reg, handleRegisterRequest, t)
	}
}

func TestUserLogin(t *testing.T) {
	for _, l := range logins {
		testUserAction(l, handleLoginRequest, t)
	}
}

func testUserAction(l testLogin, userAction func(string) (string, error), t *testing.T) {
	data := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, l.email, l.password)

	token, err := userAction(data)

	if err == nil && l.err != nil {
		t.Errorf("Error should have been received: %v", l.err)
	}

	if err != nil && l.err == nil {
		t.Errorf("Error should not have been received: %v", err)
	}

	if err != nil && l.err != nil && err.Error() != l.err.Error() {
		t.Errorf("Incorrect error received. Expected: %v, Received: %v", l.err, err)
	}

	if l.token != "" && strings.HasPrefix(token, l.token) {
		t.Errorf("Expected token not found. Expected: %s, Received: %s", l.token, token)
	}

	if l.token == "" && token != "" {
		t.Errorf("Token expected to be empty. Found: %s", token)
	}
}

var handleRegisterRequest = func(data string) (string, error) {
	return handleAuthRequest("/api/user/register", data)
}

var handleLoginRequest = func(data string) (string, error) {
	return handleAuthRequest("/api/user/login", data)
}

func handleAuthRequest(url string, data string) (string, error) {
	res, err := post(url, data)
	if err != nil {
		return "", fmt.Errorf("Unsuccessful post: %v\n", err)
	}

	if !strings.Contains(res.header.Get("Content-Type"), "application/json") {
		return "", fmt.Errorf("Incorrect header. Expected: %s, Received: %s", "application/json", res.header.Get("Content-Type"))
	}

	if len(res.body) > 0 {
		var r map[string]interface{}
		err = json.Unmarshal(res.body, &r)
		if err != nil {
			return "", err
		}

		result := handleResponse("error", r)
		if result != "" {
			return "", errors.New(result)
		}
		result = handleResponse("token", r)
		if result != "" {
			return result, nil
		}

	} else {
		fmt.Println("Empty body found when registering user.")
	}

	return "", errors.New("No expected response found.")
}

func handleResponse(t string, r map[string]interface{}) string {
	if e, ok := r[t]; ok {
		if s, ok := e.(string); ok {
			return s
		}
		return "String value not found for json error."
	}

	return ""
}
