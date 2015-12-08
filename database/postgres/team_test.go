package postgres

import (
	"testing"

	"leaguelog/model"
)

func testFindTeamByID(t *testing.T) {
	teamID := 1
	teamName := "Team 1"

	team, err := repo.FindTeamByID(teamID)
	if err != nil {
		t.Errorf("Error finding team: %v", err)
	}

	assertTeam(t, team, teamID, teamName)
}

func testCreateTeam(t *testing.T) {
	leagueID := 1
	league, err := repo.FindLeagueByID(leagueID)
	if err != nil {
		t.Errorf("Error finding league: %v", err)
	}

	team := &model.Team{
		League: league,
		Name:   "Team 3",
	}

	err = repo.CreateTeam(team)
	if err != nil {
		t.Error("Error creating team.", err)
	}

	pTeam, err := repo.FindTeamByID(team.ID)
	if err != nil {
		t.Errorf("Error finding team by id: %d", team.ID)
		t.Error(err)
	}

	assertTeam(t, pTeam, team.ID, team.Name)
}

func assertTeam(t *testing.T, team *model.Team, id int, name string) {
	if team == nil {
		t.Errorf("Error finding team id: %d", id)
	}

	if team.ID != id {
		t.Errorf("Team ids do not match. Expected: %d, Received: %d", id, team.ID)
	}

	if team.Name != name {
		t.Errorf("Team names do not match. Expected: %s, Received: %s", name, team.Name)
	}
}
