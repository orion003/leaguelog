package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"leaguelog/auth/jwt"
	"leaguelog/cmd/config"
	"leaguelog/controller"
	"leaguelog/database/postgres"
	"leaguelog/logging"
)

var conf config.Config

func main() {
	err := initializeConfig()
	if err != nil {
		fmt.Printf("Unable to initialize config: %v\n", err)
		os.Exit(1)
	}

	log := logging.NewLog15()
	c := controller.NewController(log)

	initializeRepo(c)

	fmt.Printf("Initializing root route: %s \n", conf.Routing.Root)
	r := controller.NewRouter(c, conf.Routing.Root)

	fmt.Printf("Listening on port %s\n", conf.Routing.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Routing.Port), r)
}

func initializeRepo(c *controller.Controller) {
	manager, err := postgres.NewPgManager(conf.Database.Url)
	if err != nil {
		fmt.Printf("Unable to intialize repo: %v\n", err)
		os.Exit(1)
	}

	repo := postgres.NewPgRepository(manager)

	c.SetRepository(repo)
}

func initializeTokenService(c *controller.Controller) {
	ts := jwt.InitializeJwt([]byte(conf.Auth.Key))

	c.SetTokenService(ts)
}

func initializeConfig() error {
	file, err := ioutil.ReadFile("./cmd/ll_koding/_config.json")
	if err != nil {
		return errors.New("Unable to open file: _config.json")
	}

	json.Unmarshal(file, &conf)

	return nil
}
