package postgres

import (
	"testing"
	"time"

	"leaguelog/model"
)

func testFindMostRecentSeasonByLeague(t *testing.T) {
	leagueId := 1

	league, err := repo.FindLeagueById(leagueId)
	if err != nil {
		t.Errorf("Unable to find league for season: %v", err)
	}

	season, err := repo.FindMostRecentSeasonByLeague(league)
	if err != nil {
		t.Errorf("Error finding season: %v", err)
	}

	seasonId := 1
	assertSeason(t, season, seasonId, "Test League 1 Season")
}

func testCreateSeason(t *testing.T) {
	leagueId := 1

	league, err := repo.FindLeagueById(leagueId)
	if err != nil {
		t.Errorf("Unable to find league for season: %v", err)
	}

	startDate := time.Date(2015, time.October, 6, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2016, time.April, 24, 0, 0, 0, 0, time.UTC)
	season := &model.Season{
		League:    league,
		Name:      "Test Season 2",
		StartDate: startDate,
		EndDate:   endDate,
	}

	err = repo.CreateSeason(season)
	if err != nil {
		t.Error("Error creating season.", err)
	}

	persistedSeason, err := repo.FindSeasonById(season.Id)
	if err != nil {
		t.Errorf("Error finding season by id: %d", season.Id)
		t.Error(err)
	}

	assertSeason(t, persistedSeason, season.Id, season.Name)
}

func assertSeason(t *testing.T, season *model.Season, id int, name string) {
	if season == nil {
		t.Errorf("No season found with id: %d\n", id)
	}

	if season.Id != id {
		t.Errorf("Id %d not set for season.\n", id)
	}

	if season.Name != name {
		t.Errorf("Name %s not set for season.\n", name)
	}
}
