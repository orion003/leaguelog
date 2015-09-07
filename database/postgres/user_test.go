package postgres

import (
	"testing"
)

func testCreateUser(t *testing.T) {
	truncateTables()

	user, err := createUser(userRepo)
	if err != nil {
		t.Error("Error creating user.", err)
	}

	persistedUser, err := userRepo.FindById(user.Id)
	if err != nil {
		t.Errorf("Error finding user by id: %d", user.Id)
		t.Error(err)
	}

	if user.Id != persistedUser.Id {
		t.Error("User IDs do not match.")
	}

	if user.Email != persistedUser.Email {
		t.Error("User emails do not match.")
	}
}

func testFindUserByEmail(t *testing.T) {
	truncateTables()

	user, err := createUser(userRepo)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	persistedUser, err := userRepo.FindByEmail(user.Email)
	if err != nil {
		t.Errorf("Error finding user id %d: %v", user.Id, err)
	}

	if user.Id != persistedUser.Id {
		t.Error("User IDs do not match.")
	}
}
