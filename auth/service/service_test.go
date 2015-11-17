package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockUser struct {
	id       string
	password string
}

func (u *MockUser) Exists() error {
	if u.id != "test@leaguelog.ca" {
		return errors.New(fmt.Sprintf("Authentication failed - ID is incorrect: %s", u.id))
	}
	if u.password != "password" {
		return errors.New(fmt.Sprintf("Authentication failed - Password is incorrect: %s", u.password))
	}

	return nil
}

func (u *MockUser) Save() error {
	return nil
}

func (u *MockUser) Identifier() string {
	return u.id
}

func (u *MockUser) Claims() map[string]interface{} {
	claims := make(map[string]interface{})
	claims["id"] = u.id

	return claims
}

func TestUserAuthenticateAndGenerateToken(t *testing.T) {
	user := &MockUser{
		id:       "test@leaguelog.ca",
		password: "password",
	}

	key, err := ioutil.ReadFile("testdata/hmac_key")
	if err != nil {
		t.Fatal("Unable to read HMAC key.")
	}

	j := InitializeJwt(key)

	auth := InitializeAuthentication(user, j)
	status, token := auth.Authenticate()

	if status != http.StatusOK {
		t.Errorf("Status should be %d and not %d", http.StatusOK, status)
	}

	expectedToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.mkGPonVdQeQ1nTezFMVKGHvZiAY9L1dLrLDmhcV-4gvZ4bGOm8J1jSlh5eRd-eSWrOkiZqnIHRU4i1gELq2S2A"
	if token != expectedToken {
		t.Error("Incorrect token - expected: %s, received: %s", expectedToken, token)
	}
}

func TestUserAuthenticateFailPassword(t *testing.T) {
	user := &MockUser{
		id:       "test@leaguelog.ca",
		password: "invalid_password",
	}

	key, err := ioutil.ReadFile("testdata/hmac_key")
	if err != nil {
		t.Fatal("Unable to read HMAC key.")
	}

	j := InitializeJwt(key)

	auth := InitializeAuthentication(user, j)
	status, token := auth.Authenticate()

	if status != http.StatusUnauthorized {
		t.Errorf("Status should be %d and not %d", http.StatusUnauthorized, status)
	}

	if token != "" {
		t.Error("Token should be empty because status is unauthorized.")
	}
}

func TestValidateTokenService(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.mkGPonVdQeQ1nTezFMVKGHvZiAY9L1dLrLDmhcV-4gvZ4bGOm8J1jSlh5eRd-eSWrOkiZqnIHRU4i1gELq2S2A"

	key, err := ioutil.ReadFile("testdata/hmac_key")
	if err != nil {
		t.Fatal("Unable to read HMAC key.")
	}

	j := InitializeJwt(key)
	val := InitializeValidation(j)

	err = val.ValidateToken(token)
	if err != nil {
		t.Errorf("Token should be valid (%v): %s", err, token)
	}
}

func TestInvalidateTokenService(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImludmFsaWRAbGVhZ3VlbG9nLmNhIn0.mkGPonVdQeQ1nTezFMVKGHvZiAY9L1dLrLDmhcV-4gvZ4bGOm8J1jSlh5eRd-eSWrOkiZqnIHRU4i1gELq2S2A"

	key, err := ioutil.ReadFile("testdata/hmac_key")
	if err != nil {
		t.Fatal("Unable to read HMAC key.")
	}

	j := InitializeJwt(key)
	val := InitializeValidation(j)

	err = val.ValidateToken(token)
	if err == nil {
		t.Errorf("Token should not be valid: %s", err, token)
	}
}
