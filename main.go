package main

import (
	"fmt"
	"net/http"
	"os"

	"recleague/controller"
	"recleague/database/postgres"
)

func main() {
	c := &controller.Controller{}
	initializeRepos(c)

	r := controller.NewRouter(c)

	fmt.Print("Listening on port 8000")
	http.ListenAndServe(":8000", r)
}

func initializeRepos(c *controller.Controller) {
	manager, err := postgres.NewPgManager("")
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
