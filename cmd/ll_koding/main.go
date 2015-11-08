package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"leaguelog/cmd/config"
	"leaguelog/controller"
	"leaguelog/database/postgres"
	"leaguelog/logging"
)

var conf config.Config

func main() {
	initializeConfig()

	log := logging.NewLog15()
	c := controller.NewController(log)

	initializeRepos(c)

	fmt.Printf("Initializing root route: %s \n", conf.Routing.Root)
	r := controller.NewRouter(c, conf.Routing.Root)

	fmt.Printf("Listening on port %s\n", conf.Routing.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Routing.Port), r)
}

func initializeRepos(c *controller.Controller) {
	manager, err := postgres.NewPgManager(conf.Database.Url)
	if err != nil {
		os.Exit(1)
	}

	repo := postgres.NewPgRepository(manager)

	c.SetRepository(repo)
}

func initializeConfig() error {
	file, err := ioutil.ReadFile("./_config.json")
	if err != nil {
		return errors.New("Unable to open file: _config.json")
	}

	json.Unmarshal(file, &conf)

	return nil
}
