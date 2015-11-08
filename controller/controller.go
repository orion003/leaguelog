package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"leaguelog/Godeps/_workspace/src/github.com/gorilla/mux"

	"leaguelog/logging"
	"leaguelog/model"
)

type Controller struct {
	repo model.Repository

	log logging.Logger
}

func NewController(l logging.Logger) *Controller {
	c := &Controller{
		log: l,
	}

	return c
}

func (c *Controller) SetRepository(repo model.Repository) {
	c.repo = repo
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
		err = c.repo.CreateUser(&user)
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
	leagues, err := c.repo.FindAllLeagues()
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

	league, err := c.repo.FindLeagueById(leagueId)
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
	season, err := c.repo.FindMostRecentSeasonByLeague(league)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding season for league %d: %v", league.Id, err))
	}
	standings, err := c.repo.FindAllStandingsBySeason(season)
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
	season, err := c.repo.FindMostRecentSeasonByLeague(league)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding season for league %d: %v", league.Id, err))
	}

	year, month, day := time.Now().Date()
	gameTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	games, err := c.repo.FindAllGamesAfterDateBySeason(season, &gameTime)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding games for season %d: %v", season.Id, err))
	}

	schedule := make(map[time.Time][]model.Game)
	numDates := 0
	for _, game := range games {
		startDate := time.Date(game.StartTime.Year(), game.StartTime.Month(), game.StartTime.Day(), 0, 0, 0, 0, time.UTC)
		_, ok := schedule[startDate]
		if ok == false {
			schedule[startDate] = make([]model.Game, 0, 1)
			numDates = numDates + 1
		}

		schedule[startDate] = append(schedule[startDate], game)
	}

	schedules := make([]Schedule, 0, numDates)
	for d, g := range schedule {
		s := Schedule{
			StartDate: d,
			Games:     g,
		}

		schedules = append(schedules, s)
	}

	err = c.jsonResponse(w, schedules)
	if err != nil {
		c.log.Error(fmt.Sprintf("Unable to get league games: %v", err))
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
