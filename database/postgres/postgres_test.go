package postgres

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"leaguelog/logging"

	_ "leaguelog/Godeps/_workspace/src/github.com/lib/pq"
)

var log logging.Logger = logging.NewLog15()

var repo *PgRepository

type config struct {
	Database database `json:"database"`
}

type database struct {
	URL  string `json:"url"`
	Seed string `json:"seed"`
	Test string `json:"test"`
}

var c config

func TestMain(m *testing.M) {
	err := initialize()
	if err != nil {
		log.Error(fmt.Sprintf("Unable to initialize database: %v", err))
		os.Exit(1)
	}

	url := c.Database.URL + c.Database.Test

	manager, err := NewPgManager(url)
	if err != nil {
		log.Error("Unable to initialize PgManager")
		os.Exit(1)
	}

	repo = NewPgRepository(manager)

	r := m.Run()

	err = manager.close()
	if err != nil {
		log.Error(fmt.Sprintf("Error closing database: %v", err))
		os.Exit(1)
	}

	os.Exit(r)
}

func TestLeague(t *testing.T) {
	testFindLeagueByID(t)
	testCreateLeague(t)
}

func TestSeason(t *testing.T) {
	testFindMostRecentSeasonByLeague(t)
	testCreateSeason(t)
}

func TestTeam(t *testing.T) {
	testFindTeamByID(t)
	testCreateTeam(t)
}

func TestStanding(t *testing.T) {
	testFindStandingsBySeason(t)
}

func TestGame(t *testing.T) {
	testFindAllGamesAfterDateBySeason(t)
	testCreateGame(t)
}

func TestUser(t *testing.T) {
	testFindUserByID(t)
	testFindUserByEmail(t)
	testCreateUser(t)
	testInvalidUserEmail(t)
	testDuplicateUserEmail(t)
}

//helper functions

func initialize() error {
	file, err := ioutil.ReadFile("./_config.json")
	if err != nil {
		return errors.New("Unable to open file: _config.json")
	}

	json.Unmarshal(file, &c)

	err = initializeTables()
	if err != nil {
		return err
	}

	return nil
}

func initializeTables() error {
	db, err := sql.Open("postgres", c.Database.URL+c.Database.Seed)
	if err != nil {
		return err
	}

	defer db.Close()

	log.Info(fmt.Sprintf("Dropping database: %s", c.Database.Test))
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", c.Database.Test))
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Creating database %s from template %s", c.Database.Test, c.Database.Seed))
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s TEMPLATE %s", c.Database.Test, c.Database.Seed))
	if err != nil {
		return err
	}

	return nil
}
