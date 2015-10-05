package postgres

import (
	"testing"
)

func testGetStandingBySeason(t *testing.T) {
	truncateTables()

	league, err := createLeague(leagueRepo)
	if err != nil {
		t.Error("Error creating league.", err)
	}

	season, err := createSeason(seasonRepo, league)
	if err != nil {
		t.Error("Error creating season.", err)
	}

	team, err := createTeam(teamRepo, league)
	if err != nil {
		t.Error("Error creating team.", err)
	}

	_, err = createStanding(standingRepo, season, team)
	if err != nil {
		t.Error("Error creating standing.", err)
	}

	_, err = createStanding(standingRepo, season, team)
	if err != nil {
		t.Error("Error creating standing.", err)
	}

	persistedStandings, err := standingRepo.FindAllBySeason(season)
	if err != nil {
		t.Errorf("Error finding standing by season: %d", season.Id)
		t.Error(err)
	}

	if len(persistedStandings) != 2 {
		t.Errorf("Incorrect number of standings found: %d", len(persistedStandings))
	}
}
