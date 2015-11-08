package postgres

import (
	"testing"
)

func testFindStandingsBySeason(t *testing.T) {
	seasonId := 1

	season, err := repo.FindSeasonById(seasonId)
	if err != nil {
		t.Errorf("Error finding season: %v", err)
	}

	standings, err := repo.FindAllStandingsBySeason(season)
	if err != nil {
		t.Errorf("Error finding standings: %v", err)
	}

	if len(standings) != 2 {
		t.Errorf("Wrong number of standings by season: %d", len(standings))
	}

	for _, standing := range standings {
		if standing.Id == 1 {
			if standing.Wins != 2 && standing.Losses != 0 && standing.Ties != 0 {
				t.Errorf("Standing %d has incorrect number of w-l-t: %d-%d-%d", standing.Id, standing.Wins, standing.Losses, standing.Ties)
			}
		} else if standing.Id == 2 {
			if standing.Wins != 0 && standing.Losses != 2 && standing.Ties != 0 {
				t.Errorf("Standing %d has incorrect number of w-l-t: %d-%d-%d", standing.Id, standing.Wins, standing.Losses, standing.Ties)
			}
		} else {
			t.Errorf("Expected standing not found: %d", standing.Id)
		}
	}
}
