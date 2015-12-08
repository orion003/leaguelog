package model

import (
	"errors"
)

var UserDuplicateEmail error = errors.New("user_duplicate_email")
var UserInvalidEmail error = errors.New("user_invalid_email")
var UserInvalidPassword error = errors.New("user_invalid_password")
