package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"recleague/Godeps/_workspace/src/github.com/gorilla/mux"

	"recleague/logging"
	"recleague/model"
)

type Controller struct {
	leagueRepo   model.LeagueRepository
	seasonRepo   model.SeasonRepository
	teamRepo     model.TeamRepository
	standingRepo model.StandingRepository
	gameRepo     model.GameRepository
	userRepo     model.UserRepository

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

func (c *Controller) SetStandingRepository(repo model.StandingRepository) {
	c.standingRepo = repo
}

func (c *Controller) SetGameRepository(repo model.GameRepository) {
	c.gameRepo = repo
}

func (c *Controller) SetUserRepository(repo model.UserRepository) {
	c.userRepo = repo
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.log.Info(fmt.Sprintf("Index Request: %s", r.URL.Path))
	http.ServeFile(w, r, root+"index.html")
}

func (c *Controller) AddEmail(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

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

func (c *Controller) GetLeague(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leagueId, err := strconv.Atoi(vars["leagueId"])
	if err != nil {
		c.log.Error("League ID not available: %v", err)
	}

	league, err := c.leagueRepo.FindById(leagueId)
	if err != nil {
		c.log.Error("Unable to find league: %v", err)
	}

	err = c.jsonResponse(w, league)
	if err != nil {
		c.log.Error("Unable to get league: %v", err)
	}
}

func (c *Controller) GetLeagueStandings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leagueId, err := strconv.Atoi(vars["leagueId"])

	league := &model.League{Model: model.Model{Id: leagueId}}
	season, err := c.seasonRepo.FindMostRecentByLeague(league)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding season for league %d: %v", league.Id, err))
	}
	standings, err := c.standingRepo.FindAllBySeason(season)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding standings for season %d: %v", season.Id, err))
	}

	err = c.jsonResponse(w, standings)
	if err != nil {
		c.log.Error("Unable to get league standings: %v", err)
	}
}

func (c *Controller) GetLeagueSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leagueId, err := strconv.Atoi(vars["leagueId"])

	league := &model.League{Model: model.Model{Id: leagueId}}
	season, err := c.seasonRepo.FindMostRecentByLeague(league)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding season for league %d: %v", league.Id, err))
	}
	games, err := c.gameRepo.FindUpcomingBySeason(season)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding games for season %d: %v", season.Id, err))
	}

	err = c.jsonResponse(w, games)
	if err != nil {
		c.log.Error("Unable to get league games: %v", err)
	}
}

func (c *Controller) jsonResponse(w http.ResponseWriter, v interface{}) error {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return err
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}

	return nil
}

func jsonError(err error) map[string]string {
	m := make(map[string]string)
	m["error"] = err.Error()

	return m
}
