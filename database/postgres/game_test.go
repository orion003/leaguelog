package postgres

import (
	"testing"
)

func testCreateGame(t *testing.T) {
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
