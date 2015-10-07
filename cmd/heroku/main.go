package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"recleague/cmd/config"
	"recleague/controller"
	"recleague/database/postgres"
	"recleague/logging"
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

	initializeRepos(c)

	fmt.Printf("Initializing root route: %s \n", conf.Routing.Root)
	r := controller.NewRouter(c, conf.Routing.Root)

	fmt.Printf("Listening on port %s\n", conf.Routing.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Routing.Port), r)
}

func initializeRepos(c *controller.Controller) {
	manager, err := postgres.NewPgManager(conf.Database.Url)
	if err != nil {
	    fmt.Printf("Unable to intialize repos: %v\n", err)
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
    port := os.Getenv("PORT")
    if port == "" {
        return errors.New("Unable to determine the port.")   
    }
    
    root := os.Getenv("ROOT")
    if root == "" {
        return errors.New("Unable to determine the root.")   
    }
    
    db := os.Getenv("DATABASE")
    if db == "" {
        return errors.New("Unable to determine the database.")   
    }
    
    conf = config.Config {
        Database: config.Database{Url: db,},
        Routing: config.Routing{Root: root, Port: port,},
    }

	return nil
}
