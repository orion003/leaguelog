package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"leaguelog/Godeps/_workspace/src/github.com/gorilla/mux"

	"leaguelog/auth"
	"leaguelog/logging"
	"leaguelog/model"
)

type Controller struct {
	repo  model.Repository
	token auth.TokenService

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
	leagueID, err := strconv.Atoi(vars["leagueID"])
	if err != nil {
		c.log.Error(fmt.Sprintf("League ID not available: %v", err))
	}

	league, err := c.repo.FindLeagueByID(leagueID)
	if err != nil {
		c.log.Error(fmt.Sprintf("Unable to find league: %v", err))
	}

	err = c.jsonResponse(w, league, http.StatusOK)
	if err != nil {
		c.log.Error(fmt.Sprintf("Unable to get league: %v", err))
	}
}

func (c *Controller) GetLeagueStandings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leagueID, err := strconv.Atoi(vars["leagueID"])

	league := &model.League{Model: model.Model{ID: leagueID}}
	season, err := c.repo.FindMostRecentSeasonByLeague(league)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding season for league %d: %v", league.ID, err))
	}
	standings, err := c.repo.FindAllStandingsBySeason(season)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding standings for season %d: %v", season.ID, err))
	}

	err = c.jsonResponse(w, standings, http.StatusOK)
	if err != nil {
		c.log.Error(fmt.Sprintf("Unable to get league standings: %v", err))
	}
}

func (c *Controller) GetLeagueSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leagueID, err := strconv.Atoi(vars["leagueID"])

	league := &model.League{Model: model.Model{ID: leagueID}}
	season, err := c.repo.FindMostRecentSeasonByLeague(league)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding season for league %d: %v", league.ID, err))
	}

	year, month, day := time.Now().Date()
	gameTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	games, err := c.repo.FindAllGamesAfterDateBySeason(season, &gameTime)
	if err != nil {
		c.log.Error(fmt.Sprintf("Error finding games for season %d: %v", season.ID, err))
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

	err = c.jsonResponse(w, schedules, http.StatusOK)
	if err != nil {
		c.log.Error(fmt.Sprintf("Unable to get league games: %v", err))
	}
}

func (c *Controller) jsonResponse(w http.ResponseWriter, v interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		c.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func jsonError(err error) map[string]string {
	m := make(map[string]string)
	m["error"] = err.Error()

	return m
}
