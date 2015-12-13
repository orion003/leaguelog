package model

import (
	"errors"
)

var UserDuplicateEmail = errors.New("user_duplicate_email")
var UserInvalidEmail = errors.New("user_invalid_email")
var UserInvalidPassword = errors.New("user_invalid_password")
var UserIncorrectPassword = errors.New("user_incorrect_password")
var UserUnknownEmail = errors.New("user_unknown_email")
