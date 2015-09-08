package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"recleague/config"
	"recleague/controller"
	"recleague/database/postgres"
	"recleague/logging"
)

var conf config.Config

func main() {
	initializeConfig()

	log := logging.NewLog15()
	c := controller.NewController(log)

	initializeRepos(c)

	r := controller.NewRouter(c)

	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", r)
}

func initializeRepos(c *controller.Controller) {
	manager, err := postgres.NewPgManager(conf.Database.Url)
	if err != nil {
		os.Exit(1)
	}

	leagueRepo := postgres.NewPgLeagueRepository(manager)
	seasonRepo := postgres.NewPgSeasonRepository(manager)
	teamRepo := postgres.NewPgTeamRepository(manager)
	gameRepo := postgres.NewPgGameRepository(manager)
	userRepo := postgres.NewPgUserRepository(manager)

	c.SetLeagueRepository(leagueRepo)
	c.SetSeasonRepository(seasonRepo)
	c.SetTeamRepository(teamRepo)
	c.SetGameRepository(gameRepo)
	c.SetUserRepository(userRepo)
}

func initializeConfig() error {
	file, err := ioutil.ReadFile("./config/_config.json")
	if err != nil {
		return errors.New("Unable to open file: _config.json")
	}

	json.Unmarshal(file, &conf)

	return nil
}
