package postgres

import (
	"testing"

	"leaguelog/model"
)

func testFindUserById(t *testing.T) {
	id := 1
	user, err := repo.FindUserById(id)

	if err != nil {
		t.Errorf("Error finding user by id: %v\n", err)
	}

	assertUser(t, user, id, "test@leaguelog.ca")
}

func testFindUserByEmail(t *testing.T) {
	email := "test@leaguelog.ca"
	user, err := repo.FindUserByEmail(email)

	if err != nil {
		t.Errorf("Error finding user by id: %v\n", err)
	}

	assertUser(t, user, 1, email)
}

func testCreateUser(t *testing.T) {
	user := &model.User{
		Email: "test2@leaguelog.ca",
	}

	err := repo.CreateUser(user)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	persistedUser, err := repo.FindUserById(user.Id)
	if err != nil {
		t.Errorf("Error finding user: %v", err)
	}

	assertUser(t, persistedUser, user.Id, user.Email)
}

func testInvalidUserEmail(t *testing.T) {
	user := &model.User{
		Email: "test_invalid_email",
	}

	err := repo.CreateUser(user)
	if err == nil {
		t.Errorf("User email should be invalid: %s", user.Email)
	}

	if err != model.UserInvalidEmail {
		t.Errorf("Error type should be %s", "UserInvalidEmail")
	}
}

func testDuplicateUserEmail(t *testing.T) {
	user := &model.User{
		Email: "test@leaguelog.ca",
	}

	err := repo.CreateUser(user)
	if err == nil {
		t.Errorf("User email should be duplicate: %s", user.Email)
	}

	if err != model.UserDuplicateEmail {
		t.Errorf("Error type should be %s", "UserDuplicateEmail")
	}
}

func assertUser(t *testing.T, user *model.User, id int, email string) {
	if user == nil {
		t.Errorf("Unable to find user with id: %d", id)
	}

	if user.Id != id {
		t.Errorf("User ids do not match. Expected: %d, Received: %d", id, user.Id)
	}

	if user.Email != email {
		t.Errorf("User emails do not match. Expected: %s, Received: %s", email, user.Email)
	}
}
