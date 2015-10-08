package postgres

import (
	"time"

	"leaguelog.ca/database/marshal"
	"leaguelog.ca/model"
)

func NewPgSeasonRepository(manager *PgManager) *PgSeasonRepository {
	repo := &PgSeasonRepository{
		manager: manager,
	}

	return repo
}

func (repo *PgSeasonRepository) Create(season *model.Season) error {
	err := season.Validate(repo)
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO season(league_id, name, start_date, end_date, created, modified) 
	    VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		season.League.Id, season.Name, season.StartDate, season.EndDate, t, t).Scan(&id)

	if err != nil {
		return err
	}

	season.Id = id
	season.Created = t
	season.Modified = t

	return nil
}

func (repo *PgSeasonRepository) FindById(id int) (*model.Season, error) {
	row := repo.manager.db.QueryRow(`SELECT s.id, s.league_id, s.name, s.start_date, s.end_date, s.created, s.modified 
        FROM season s
        WHERE s.id = $1`, id)

	season, err := marshal.Season(row)
	if err != nil {
		return &model.Season{}, err
	}

	return season, nil
}

func (repo *PgSeasonRepository) FindMostRecentByLeague(league *model.League) (*model.Season, error) {
	row := repo.manager.db.QueryRow(`SELECT s.id, s.league_id, s.name, s.start_date, s.end_date, s.created, s.modified
        FROM season s
        WHERE s.league_id = $1
        ORDER BY s.end_date DESC
        LIMIT 1`, league.Id)

	season, err := marshal.Season(row)
	if err != nil {
		return &model.Season{}, err
	}

	return season, err
}
