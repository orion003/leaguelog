package postgres

import (
	"testing"
)

func testCreateLeague(t *testing.T) {
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
