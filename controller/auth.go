package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"leaguelog/auth"
	"leaguelog/model"
)

type AuthUser struct {
	controller *Controller
	user       *model.User
}

func (au *AuthUser) Save() error {
	repo := au.controller.repo

	salt := auth.GenerateRandomSalt(22)
	hash, err := auth.GenerateHashedPassword(salt, au.user.Password)
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
	repo := au.controller.repo
	user, err := repo.FindUserByEmail(au.user.Email)
	if err != nil {
		au.controller.log.Error(fmt.Sprintf("Unable to find user (%s) by email: %v", au.user.Email, err))
		return model.UserUnknownEmail
	}
	if user == nil {
		au.controller.log.Error(fmt.Sprintf("Email not found: %v", au.user.Email))
		return model.UserUnknownEmail
	}

	err = auth.CompareHashAndPassword(user.Password, user.Salt, au.user.Password)
	if err != nil {
		return model.UserIncorrectPassword
	}

	return nil
}

func (au *AuthUser) Claims() map[string]interface{} {
	claims := make(map[string]interface{})
	claims["iss"] = "leaguelog.ca"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["uid"] = au.user.Email

	return claims
}

func (c *Controller) SetTokenService(t auth.TokenService) {
	c.token = t
}

func (c *Controller) UserRegister(w http.ResponseWriter, r *http.Request) {
	c.userControllerHandler(w, r, register)
}

func (c *Controller) UserLogin(w http.ResponseWriter, r *http.Request) {
	c.userControllerHandler(w, r, login)
}

func (c *Controller) userControllerHandler(w http.ResponseWriter, r *http.Request, fn func(c *Controller, u *model.User) (string, error)) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		c.log.Error(fmt.Sprintf("Unable to decode user JSON: %v", err))
		w.WriteHeader(http.StatusNotAcceptable)
	}

	token, err := fn(c, &user)
	if err != nil {
		handleAuthenticationError(w, err)
	} else {
		handleTokenResponse(w, token)
	}
}

var register = func(controller *Controller, user *model.User) (string, error) {
	us := AuthUser{user: user, controller: controller}
	tokenService := controller.token

	auth := auth.InitializeAuthentication(&us, tokenService)

	token, err := auth.Register()
	if err != nil {
		return "", err
	}

	return token, nil
}

var login = func(controller *Controller, user *model.User) (string, error) {
	us := AuthUser{user: user, controller: controller}
	tokenService := controller.token

	auth := auth.InitializeAuthentication(&us, tokenService)

	token, err := auth.Authenticate()
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

func handleAuthenticationError(w http.ResponseWriter, err error) {
	fmt.Printf("RegisterUser error: %v\n", err)
	w.WriteHeader(http.StatusNotAcceptable)

	m := make(map[string]string)
	m["error"] = err.Error()

	if e := json.NewEncoder(w).Encode(m); e != nil {
		fmt.Printf("Error encoding error: %v\n", e)
	}
}
