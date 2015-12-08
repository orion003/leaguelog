package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"leaguelog/auth/service/Godeps/_workspace/src/golang.org/x/crypto/bcrypt"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type UserService interface {
	Exists() error
	Save() error
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

func (auth *Authenticator) Register() (string, error) {
	err := auth.user.Save()
	if err != nil {
		fmt.Printf("User not saved: %v\n", err)
		return "", err
	}

	token, err := generateToken(auth)
	if err != nil {
		fmt.Printf("Unable to generate token: %v\n", err)
		return "", err
	}

	return token, nil
}

func (auth *Authenticator) Authenticate() (string, error) {
	err := auth.user.Exists()
	if err != nil {
		fmt.Printf("User not found: %v\n", err)
		return "", err
	}

	token, err := generateToken(auth)
	if err != nil {
		fmt.Printf("Unable to generate token: %v\n", err)
		return "", err
	}

	return token, nil
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

func GenerateRandomSalt(n int) string {
	rand.Seed(time.Now().UnixNano())
	length := len(characters)

	b := make([]byte, n)
	for i := range b {
		b[i] = characters[rand.Intn(length)]
	}

	return string(b)
}

func GenerateHashedPassword(salt string, password string) (string, error) {
	combined := salt + password
	hashed, err := bcrypt.GenerateFromPassword([]byte(combined), 11)

	return string(hashed), err
}

func CompareHashAndPassword(hash string, salt string, password string) error {
	combined := salt + password

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(combined))
}
