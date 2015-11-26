package controller

import (
	"fmt"
	"testing"
)

func testUserRegister(t *testing.T) {
    email := "test_register@leaguelog.ca"
    password := "test_password"
    
	data := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)
	
	err := addEmail(email)
	if err != nil {
		t.Errorf("Unable to add email: %s - %v", email, err)
	}
}