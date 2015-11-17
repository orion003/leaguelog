package service

import (
	"errors"
	"fmt"
	"net/http"
)

type UserService interface {
	Exists() error
	Save() error
	Identifier() string
	Claims() map[string]interface{}
}

type TokenService interface {
	Generate(claims map[string]interface{}) string
	Validate(token string) error
}

type Authenticator struct {
	user  UserService
	token TokenService
}

type Validator struct {
	token TokenService
}

func InitializeAuthentication(u UserService, t TokenService) *Authenticator {
	auth := new(Authenticator)
	auth.user = u
	auth.token = t

	return auth
}

func (auth *Authenticator) Authenticate() (int, string) {
	err := auth.user.Exists()
	if err != nil {
		fmt.Printf("User not found: %v\n", err)
		return http.StatusUnauthorized, ""
	}

	token, err := generateToken(auth)
	if err != nil {
		fmt.Printf("Unable to generate token: %v\n", err)
		return http.StatusUnauthorized, ""
	}

	return http.StatusOK, token
}

func generateToken(auth *Authenticator) (string, error) {
	tokenService := auth.token

	token := tokenService.Generate(auth.user.Claims())
	if token == "" {
		return "", errors.New("Token has not been generated.")
	}

	return token, nil
}

func InitializeValidation(t TokenService) *Validator {
	v := &Validator{
		token: t,
	}

	return v
}

func (val *Validator) ValidateToken(token string) error {
	return val.token.Validate(token)
}
