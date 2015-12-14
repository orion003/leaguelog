package auth

import (
	"errors"
	"fmt"
	"testing"

	"leaguelog/auth/jwt"
)

type MockUser struct {
	id       string
	password string
}

var hmac = []byte("579760E50509F2F28324421C7509741F5BF03B9158161076B3C6B39FB028D9E2C251490A3F8BD1F59728259A673668CAFEB49C9E9499F8386B147D7260B6937A")

func (u *MockUser) Exists() error {
	if u.id != "test@leaguelog.ca" {
		return fmt.Errorf("Authentication failed - ID is incorrect: %s", u.id)
	}
	if u.password != "password" {
		return fmt.Errorf("Authentication failed - Password is incorrect: %s", u.password)
	}

	return nil
}

func (u *MockUser) Save() error {
	if u.id == "" || u.password == "" {
		return errors.New("User save failed. Empty username or password")
	}

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

func TestGeneratePassword(t *testing.T) {
	salt := "8YrYmqwWonIIhpUjHaSHtB"
	password := "password"

	hashed, err := GenerateHashedPassword(salt, password)
	if err != nil {
		t.Errorf("There should be no error hashing password: %v", err)
	}

	err = CompareHashAndPassword(hashed, salt, password)
	if err != nil {
		t.Errorf("Password and hash do not match: %v", err)
	}
}

func TestUserRegisterAndGenerateToken(t *testing.T) {
	user := &MockUser{
		id:       "test@leaguelog.ca",
		password: "password",
	}

	j := jwt.InitializeJwt(hmac)

	auth := InitializeAuthentication(user, j)
	token, err := auth.Register()

	if err != nil {
		t.Errorf("No error should be returned when registering: %v", err)
	}

	expectedToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.Xx8WqprVmafi1ZaoMWiL43dC5H-Im0hTtJydprrFGWWndgeTQB91aWSan37NrIBP_rBL06Axv-MO0WJNji70kw"
	if token != expectedToken {
		t.Errorf("Incorrect token - expected: %s, received: %s", expectedToken, token)
	}
}

func TestUserAuthenticateAndGenerateToken(t *testing.T) {
	user := &MockUser{
		id:       "test@leaguelog.ca",
		password: "password",
	}

	j := jwt.InitializeJwt(hmac)

	auth := InitializeAuthentication(user, j)
	token, err := auth.Authenticate()

	if err != nil {
		t.Errorf("No error should be returned when registering: %v", err)
	}

	expectedToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.Xx8WqprVmafi1ZaoMWiL43dC5H-Im0hTtJydprrFGWWndgeTQB91aWSan37NrIBP_rBL06Axv-MO0WJNji70kw"
	if token != expectedToken {
		t.Errorf("Incorrect token - expected: %s, received: %s", expectedToken, token)
	}
}

func TestUserAuthenticateFailPassword(t *testing.T) {
	user := &MockUser{
		id:       "test@leaguelog.ca",
		password: "invalid_password",
	}

	j := jwt.InitializeJwt(hmac)

	auth := InitializeAuthentication(user, j)
	token, err := auth.Authenticate()

	if err == nil {
		t.Errorf("Error should be returned on failed authentication.")
	}

	if token != "" {
		t.Error("Token should be empty because status is unauthorized.")
	}
}

func TestValidateTokenService(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.Xx8WqprVmafi1ZaoMWiL43dC5H-Im0hTtJydprrFGWWndgeTQB91aWSan37NrIBP_rBL06Axv-MO0WJNji70kw"

	j := jwt.InitializeJwt(hmac)
	val := InitializeValidation(j)

	err := val.ValidateToken(token)
	if err != nil {
		t.Errorf("Token should be valid (%v): %s", err, token)
	}
}

func TestInvalidateTokenService(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImludmFsaWRAbGVhZ3VlbG9nLmNhIn0.mkGPonVdQeQ1nTezFMVKGHvZiAY9L1dLrLDmhcV-4gvZ4bGOm8J1jSlh5eRd-eSWrOkiZqnIHRU4i1gELq2S2A"

	j := jwt.InitializeJwt(hmac)
	val := InitializeValidation(j)

	err := val.ValidateToken(token)
	if err == nil {
		t.Errorf("Token should not be valid (%s): %s", token, err)
	}
}
