package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"leaguelog/model"
)

func testUserRegister(t *testing.T) {
	email := "test_register@leaguelog.ca"
	password := "test_password"

	data := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)

	err := addUser(data)
	if err != nil {
		t.Errorf("Unable to add user: %s - %v", email, err)
	}
}

func testDuplicateEmailRegister(t *testing.T) {
	email := "test@leaguelog.ca"
	password := "duplicate_password"

	data := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)

	err := addUser(data)
	if err == nil {
		t.Errorf("Duplicate email not allowed: %s", email)
	}

	if err.Error() != model.UserDuplicateEmail.Error() {
		t.Errorf("Should be duplicate email error: %s", err)
	}
}

func testInvalidEmailRegister(t *testing.T) {
	email := "test_invalid"
	password := "invalid_password"

	data := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)

	err := addUser(data)
	if err == nil {
		t.Errorf("Invalid email not allowed: %s", email)
	}

	if err.Error() != model.UserInvalidEmail.Error() {
		t.Errorf("Should be invalid email error: %s", err)
	}
}

func testMissingPasswordRegister(t *testing.T) {
	email := "missing_password@leaguelog.ca"

	data := fmt.Sprintf(`{"email": "%s"}`, email)

	err := addUser(data)
	if err == nil {
		t.Error("Invalid password not allowed.")
	}

	if err.Error() != model.UserInvalidPassword.Error() {
		t.Errorf("Should be invalid password error: %s", err)
	}
}

func addUser(data string) error {
	body, err := post("/api/register", data)
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
