package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"leaguelog/auth/service"
	"leaguelog/model"
)

type AuthUser struct {
	controller *Controller
	user       *model.User
}

func (au *AuthUser) Save() error {
	repo := au.controller.repo

	salt := service.GenerateRandomSalt(22)
	hash, err := service.GenerateHashedPassword(salt, au.user.Password)
	if err != nil {
		return fmt.Errorf("Unable to generate hashed password: %v", err)
	}

	err = au.user.Validate(repo)
	if err != nil {
		return err
	}

	user := &model.User{
		Email:    au.user.Email,
		Password: hash,
		Salt:     salt,
	}

	return repo.CreateUser(user)
}

func (au *AuthUser) Exists() error {
	return errors.New("User does not exist.")
}

func (au *AuthUser) Claims() map[string]interface{} {
	return make(map[string]interface{})
}

func (c *Controller) SetTokenService(t service.TokenService) {
	c.token = t
}

func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user model.User
	err := decoder.Decode(&user)

	if err != nil {
		c.log.Error(fmt.Sprintf("Unable to decode user JSON: %v", err))
		w.WriteHeader(http.StatusNotAcceptable)
	}

	token, err := register(c, &user)
	if err != nil {
		handleRegisterError(w, err)
	} else {
		handleTokenResponse(w, token)
	}
}

func register(controller *Controller, user *model.User) (string, error) {
	us := AuthUser{user: user, controller: controller}
	tokenService := controller.token

	auth := service.InitializeAuthentication(&us, tokenService)

	token, err := auth.Register()
	if err != nil {
		return "", err
	}

	return token, nil
}

func handleTokenResponse(w http.ResponseWriter, token string) {
	fmt.Println("Generating token.")
	w.WriteHeader(http.StatusCreated)

	t := make(map[string]string)
	t["token"] = token

	if e := json.NewEncoder(w).Encode(t); e != nil {
		fmt.Printf("Error encoding token: %v\n", e)
	}
}

func handleRegisterError(w http.ResponseWriter, err error) {
	fmt.Printf("RegisterUser error: %v\n", err)
	w.WriteHeader(http.StatusNotAcceptable)

	m := make(map[string]string)
	m["error"] = err.Error()

	if e := json.NewEncoder(w).Encode(m); e != nil {
		fmt.Printf("Error encoding error: %v\n", e)
	}
}
