package postgres

import (
	"testing"
	"time"

	"leaguelog/model"
)

func testFindAllGamesAfterDateBySeason(t *testing.T) {
	seasonID := 1
	season, err := repo.FindSeasonByID(seasonID)
	if err != nil {
		t.Errorf("Error finding season: %v", err)
	}

	after := time.Date(2015, time.October, 2, 12, 0, 0, 0, time.UTC)

	games, err := repo.FindAllGamesAfterDateBySeason(season, &after)
	if err != nil {
		t.Errorf("Error finding games: %v", err)
	}

	if len(games) != 1 {
		t.Errorf("Wrong number of games found: %d", len(games))
	}
}

func testCreateGame(t *testing.T) {
	seasonID := 1
	season, err := repo.FindSeasonByID(seasonID)
	if err != nil {
		t.Errorf("Unable to find season for game: %v", err)
	}

	homeTeamID := 1
	homeTeam, err := repo.FindTeamByID(homeTeamID)
	if err != nil {
		t.Errorf("Unable to find home team for game: %v", err)
	}

	awayTeamID := 2
	awayTeam, err := repo.FindTeamByID(awayTeamID)
	if err != nil {
		t.Errorf("Unable to find away team for game: %v", err)
	}

	startTime := time.Date(2015, time.October, 5, 21, 30, 0, 0, time.UTC)
	game := &model.Game{
		Season:    season,
		HomeTeam:  homeTeam,
		AwayTeam:  awayTeam,
		StartTime: startTime,
	}

	err = repo.CreateGame(game)
	if err != nil {
		t.Error("Error creating game.", err)
	}

	after := time.Date(2015, time.October, 2, 12, 0, 0, 0, time.UTC)
	games, err := repo.FindAllGamesAfterDateBySeason(season, &after)
	if err != nil {
		t.Errorf("Error finding games: %v", err)
	}

	if len(games) != 2 {
		t.Errorf("Wrong number of games found: %d", len(games))
	}
}
