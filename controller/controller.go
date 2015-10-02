package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"recleague/logging"
	"recleague/model"
)

type Controller struct {
	leagueRepo model.LeagueRepository
	seasonRepo model.SeasonRepository
	teamRepo   model.TeamRepository
	gameRepo   model.GameRepository
	userRepo   model.UserRepository

	log logging.Logger
}

func NewController(l logging.Logger) *Controller {
	c := &Controller{
		log: l,
	}

	return c
}

func (c *Controller) SetLeagueRepository(repo model.LeagueRepository) {
	c.leagueRepo = repo
}

func (c *Controller) SetSeasonRepository(repo model.SeasonRepository) {
	c.seasonRepo = repo
}

func (c *Controller) SetTeamRepository(repo model.TeamRepository) {
	c.teamRepo = repo
}

func (c *Controller) SetGameRepository(repo model.GameRepository) {
	c.gameRepo = repo
}

func (c *Controller) SetUserRepository(repo model.UserRepository) {
	c.userRepo = repo
}

func (c *Controller) AddEmail(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		c.log.Error(fmt.Sprintf("Unable to decode user email JSON: %v", err))
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		err = c.userRepo.Create(&user)
		if err != nil {
			c.log.Error(fmt.Sprintf("AddEmail error: %v", err))
			w.WriteHeader(http.StatusNotAcceptable)

			if e := json.NewEncoder(w).Encode(jsonError(err)); e != nil {
				c.log.Error(e.Error())
			}
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func (c *Controller) GetLeagues(w http.ResponseWriter, r *http.Request) {
	leagues, err := c.leagueRepo.FindAll()
	if err != nil {
		c.log.Error(err.Error())
	}
	
	c.log.Info(fmt.Sprintf("Leagues found: %d", len(leagues)))

	if err = json.NewEncoder(w).Encode(leagues); err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
}

func jsonError(err error) map[string]string {
	m := make(map[string]string)
	m["error"] = err.Error()

	return m
}
