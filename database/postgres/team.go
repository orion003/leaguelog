package postgres

import (
	"time"

	"leaguelog.ca/database/marshal"
	"leaguelog.ca/model"
)

func NewPgTeamRepository(manager *PgManager) *PgTeamRepository {
	repo := &PgTeamRepository{
		manager: manager,
	}

	return repo
}

func (repo *PgTeamRepository) Create(team *model.Team) error {
	err := team.Validate(repo)
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO team(league_id, name, created, modified) 
	    VALUES($1, $2, $3, $4) RETURNING id`,
		team.League.Id, team.Name, t, t).Scan(&id)

	if err != nil {
		return err
	}

	team.Id = id
	team.Created = t
	team.Modified = t

	return nil
}

func (repo *PgTeamRepository) FindById(id int) (*model.Team, error) {
	row := repo.manager.db.QueryRow(`
	    SELECT 
	        t.id, t.league_id, t.name, t.created, t.modified
        FROM team t
        WHERE t.id = $1`, id)

	team, err := marshal.Team(row)
	if err != nil {
		return &model.Team{}, err
	}

	return team, nil
}
