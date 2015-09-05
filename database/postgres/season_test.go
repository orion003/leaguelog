package postgres

import (
	"testing"
)

func testCreateSeason(t *testing.T) {
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
