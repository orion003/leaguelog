package postgres

import (
	"testing"

	"leaguelog/model"
)

func testFindTeamById(t *testing.T) {
	teamId := 1
	teamName := "Team 1"

	team, err := repo.FindTeamById(teamId)
	if err != nil {
		t.Errorf("Error finding team: %v", err)
	}

	assertTeam(t, team, teamId, teamName)
}

func testCreateTeam(t *testing.T) {
	leagueId := 1
	league, err := repo.FindLeagueById(leagueId)
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

	pTeam, err := repo.FindTeamById(team.Id)
	if err != nil {
		t.Errorf("Error finding team by id: %d", team.Id)
		t.Error(err)
	}

	assertTeam(t, pTeam, team.Id, team.Name)
}

func assertTeam(t *testing.T, team *model.Team, id int, name string) {
	if team == nil {
		t.Errorf("Error finding team id: %d", id)
	}

	if team.Id != id {
		t.Errorf("Team ids do not match. Expected: %d, Received: %d", id, team.Id)
	}

	if team.Name != name {
		t.Errorf("Team names do not match. Expected: %s, Received: %s", name, team.Name)
	}
}
