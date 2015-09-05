package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	
	"recleague/database/postgres"
	"recleague/server"
)

const root string = "../web/angular/"

func main() {
    repos := initializeRepos()
    
	r := server.NewRouter()

	r.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir(root+"app/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(root+"assets/"))))

	fmt.Print("Listening on port 8000")
	http.ListenAndServe(":8000", r)
}

func initializeRepos() map[string]interface {
	manager, err := postgres.NewPgManager(url)
	if err != nil {
		log.Error("Unable to initialize PgManager")
		os.Exit(1)
	}

	leagueRepo := &postgres.PgLeagueRepository{
		manager: manager,
	}
	seasonRepo := &postgres.PgSeasonRepository{
		manager: manager,
	}
	teamRepo := &postgres.PgTeamRepository{
		manager: manager,
	}
	gameRepo := &postgres.PgGameRepository{
		manager: manager,
	}
	userRepo := &postgres.PgUserRepository{
		manager: manager,
	}
	
	repos := make(map[string]interface))
	repos["league"] = leagueRepo
	repos["season"] = seasonRepo
	repos["team"] = teamRepo
	repos["game"] = gameRepo
	repos["user"] = userRepo
	
	return repos
}