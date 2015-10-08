package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"recleague/cmd/config"
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

	leagueRepo := postgres.NewPgLeagueRepository(manager)
	seasonRepo := postgres.NewPgSeasonRepository(manager)
	teamRepo := postgres.NewPgTeamRepository(manager)
	gameRepo := postgres.NewPgGameRepository(manager)
	standingRepo := postgres.NewPgStandingRepository(manager)
	userRepo := postgres.NewPgUserRepository(manager)

	c.SetLeagueRepository(leagueRepo)
	c.SetSeasonRepository(seasonRepo)
	c.SetTeamRepository(teamRepo)
	c.SetGameRepository(gameRepo)
	c.SetStandingRepository(standingRepo)
	c.SetUserRepository(userRepo)
}

func initializeConfig() error {
	file, err := ioutil.ReadFile("./_config.json")
	if err != nil {
		return errors.New("Unable to open file: _config.json")
	}

	json.Unmarshal(file, &conf)

	return nil
}
