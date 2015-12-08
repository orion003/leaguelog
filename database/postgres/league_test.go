package postgres

import (
	"testing"

	"leaguelog/model"
)

func testFindLeagueByID(t *testing.T) {
	id := 1
	league, err := repo.FindLeagueByID(id)

	if err != nil {
		t.Errorf("Error finding league by id: %v\n", err)
	}

	assertLeague(t, league, id, "Test League 1", "hockey")
}

func testCreateLeague(t *testing.T) {
	league := &model.League{
		Name:  "Test League 2",
		Sport: "hockey",
	}

	err := repo.CreateLeague(league)
	if err != nil {
		t.Errorf("Error creating league: %v\n", err)
	}

	persistedLeague, err := repo.FindLeagueByID(league.ID)
	if err != nil {
		t.Errorf("Error finding league by id: %d\n", league.ID)
		t.Error(err)
	}

	assertLeague(t, persistedLeague, league.ID, league.Name, league.Sport)
}

func assertLeague(t *testing.T, league *model.League, id int, name string, sport string) {
	if league == nil {
		t.Errorf("No league found with id: %d\n", id)
	}

	if league.ID != id {
		t.Errorf("Id %d not set for league.\n", id)
	}

	if league.Name != name {
		t.Errorf("Name %s not set for league.\n", name)
	}

	if league.Sport != sport {
		t.Errorf("Sport %s not set for league.\n", sport)
	}
}
