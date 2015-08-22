package postgres

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"recleague/model"

	_ "github.com/lib/pq"
)

var leagueRepo *PgLeagueRepository
var seasonRepo *PgSeasonRepository
var teamRepo *PgTeamRepository
var gameRepo *PgGameRepository

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
		os.Exit(1)
	}

	url := c.Database.Url

	manager, err := NewPgManager(url)
	if err != nil {
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

	r := m.Run()

	manager.close()

	os.Exit(r)
}

//tests for league creation

func TestCreateLeague(t *testing.T) {
	truncateTables()

	league, err := createLeague(leagueRepo)
	if err != nil {
		t.Error("Error creating league.", err)
	}

	persistedLeague, err := leagueRepo.FindById(league.Id)
	if err != nil {
		t.Errorf("Error finding league by id: %d", league.Id)
		t.Error(err)
	}

	if league.Id != persistedLeague.Id {
		t.Error("Leagues do not match.")
	}
}

func TestCreateSeason(t *testing.T) {
	truncateTables()

	league, err := createLeague(leagueRepo)
	if err != nil {
		t.Error("Error creating league.", err)
	}

	season, err := createSeason(seasonRepo, league)
	if err != nil {
		t.Error("Error creating season.", err)
	}

	persistedSeason, err := seasonRepo.FindById(season.Id)
	if err != nil {
		t.Errorf("Error finding season by id: %d", season.Id)
		t.Error(err)
	}
	if persistedSeason.League == nil {
		t.Error("League should never be nil")
	}

	if season.Id != persistedSeason.Id {
		t.Error("Seasons do not match.")
	}

	if season.League.Id != persistedSeason.League.Id {
		t.Error("Season leagues do not match")
	}
}

//tests for team creation

func TestCreateTeam(t *testing.T) {
	truncateTables()

	league, err := createLeague(leagueRepo)
	if err != nil {
		t.Error("Error creating league.", err)
	}

	team, err := createTeam(teamRepo, league)
	if err != nil {
		t.Log("Error creating team.")
		t.Error(err)
	}

	persistedTeam, err := teamRepo.FindById(team.Id)
	if err != nil {
		t.Logf("Error finding team by id: %d", team.Id)
		t.Error(err)
	}
	if persistedTeam.League == nil {
		t.Error("League should never be nil")
	}

	if team.Id != persistedTeam.Id {
		t.Error("Seasons do not match.")
	}

	if team.League.Id != persistedTeam.League.Id {
		t.Error("Season leagues do not match")
	}
}

//tests for game creation

func TestCreateGame(t *testing.T) {
	truncateTables()

	league, err := createLeague(leagueRepo)
	if err != nil {
		t.Error("Error creating league.", err)
	}

	season, err := createSeason(seasonRepo, league)
	if err != nil {
		t.Error(err)
	}

	hometeam, err := createTeam(teamRepo, league)
	if err != nil {
		t.Log("Error creating home team.")
		t.Error(err)
	}
	hometeam.Name = "Home Team"

	awayteam, err := createTeam(teamRepo, league)
	if err != nil {
		t.Log("Error creating away team.")
		t.Error(err)
	}
	awayteam.Name = "Away Team"

	game, err := createGame(gameRepo, season, hometeam, awayteam)
	if err != nil {
		t.Log("Error creating game.")
		t.Error(err)
	}

	persistedGame, err := gameRepo.FindById(game.Id)
	if err != nil {
		t.Logf("Error finding game by id: %d", game.Id)
		t.Error(err)
	}
	if persistedGame.Season == nil {
		t.Error("Season should never be nil")
	}
	if persistedGame.Home_team == nil {
		t.Error("Home team should never be nil")
	}
	if persistedGame.Away_team == nil {
		t.Error("Away team should never be nil")
	}

	if game.Id != persistedGame.Id {
		t.Error("Games do not match.")
	}
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

	_, err = db.Exec("TRUNCATE league, season, team, game RESTART IDENTITY")
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
