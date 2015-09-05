package postgres

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"recleague/logging"
	"recleague/model"

	_ "github.com/lib/pq"
)

var log logging.Logger = logging.NewLog15()

var leagueRepo *PgLeagueRepository
var seasonRepo *PgSeasonRepository
var teamRepo *PgTeamRepository
var gameRepo *PgGameRepository
var userRepo *PgUserRepository

type config struct {
	Database database `json:database`
}

type database struct {
	Url string `json:url`
}

var c config

func TestMain(m *testing.M) {
	err := initialize()
	if err != nil {
		log.Error("Unable to initialize database")
		os.Exit(1)
	}

	url := c.Database.Url

	manager, err := NewPgManager(url)
	if err != nil {
		log.Error("Unable to initialize PgManager")
		os.Exit(1)
	}

	leagueRepo = &PgLeagueRepository{
		manager: manager,
	}
	seasonRepo = &PgSeasonRepository{
		manager: manager,
	}
	teamRepo = &PgTeamRepository{
		manager: manager,
	}
	gameRepo = &PgGameRepository{
		manager: manager,
	}
	userRepo = &PgUserRepository{
		manager: manager,
	}

	r := m.Run()

	manager.close()

	os.Exit(r)
}

func TestLeague(t *testing.T) {
	testCreateLeague(t)
}

func TestSeason(t *testing.T) {
	testCreateSeason(t)
}

func TestTeam(t *testing.T) {
	testCreateTeam(t)
}

func TestGame(t *testing.T) {
	testCreateGame(t)
}

func TestUser(t *testing.T) {
	testCreateUser(t)
}

//helper functions

func initialize() error {
	file, err := ioutil.ReadFile("./_config.json")
	if err != nil {
		return errors.New("Unable to open file: _config.json")
	}

	json.Unmarshal(file, &c)

	return nil
}

func truncateTables() error {
	db, err := sql.Open("postgres", c.Database.Url)
	if err != nil {
		return err
	}

	_, err = db.Exec("TRUNCATE league, season, team, game, user0 RESTART IDENTITY")
	if err != nil {
		return err
	}

	return nil
}

func createLeague(repo *PgLeagueRepository) (*model.League, error) {
	league := &model.League{
		Name:  "Test League",
		Sport: "Hockey",
	}

	err := repo.Create(league)
	if err != nil {
		return nil, err
	}

	return league, nil
}

func createSeason(repo *PgSeasonRepository, league *model.League) (*model.Season, error) {
	startDate := time.Date(2015, time.October, 6, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2016, time.April, 24, 0, 0, 0, 0, time.UTC)
	season := &model.Season{
		League:     league,
		Name:       "Test Season",
		Start_date: startDate,
		End_date:   endDate,
	}

	err := repo.Create(season)
	if err != nil {
		return nil, err
	}

	return season, nil
}

func createTeam(repo *PgTeamRepository, league *model.League) (*model.Team, error) {
	team := &model.Team{
		League: league,
		Name:   "Test Team",
	}

	err := repo.Create(team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func createGame(repo *PgGameRepository, season *model.Season, hometeam *model.Team, awayteam *model.Team) (*model.Game, error) {
	t := time.Date(2015, time.October, 6, 8, 30, 0, 0, time.UTC)
	game := &model.Game{
		Season:     season,
		Start_time: t,
		Home_team:  hometeam,
		Away_team:  awayteam,
	}

	err := repo.Create(game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func createUser(repo *PgUserRepository) (*model.User, error) {
	user := &model.User{
		Email: "parrottr3@gmail.com",
	}

	err := repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
