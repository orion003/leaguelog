package postgres

import (
	"testing"
)

func testCreateTeam(t *testing.T) {
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
