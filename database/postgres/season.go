package postgres

import (
	"time"

	"leaguelog/database/marshal"
	"leaguelog/model"
)

func (repo *PgRepository) CreateSeason(season *model.Season) error {
	err := season.Validate(repo)
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO season(league_id, name, start_date, end_date, created, modified)
	    VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		season.League.ID, season.Name, season.StartDate, season.EndDate, t, t).Scan(&id)

	if err != nil {
		return err
	}

	season.ID = id
	season.Created = t
	season.Modified = t

	return nil
}

func (repo *PgRepository) FindSeasonByID(id int) (*model.Season, error) {
	row := repo.manager.db.QueryRow(`SELECT s.id, s.league_id, s.name, s.start_date, s.end_date, s.created, s.modified
        FROM season s
        WHERE s.id = $1`, id)

	season, err := marshal.Season(row)
	if err != nil {
		return &model.Season{}, err
	}

	return season, nil
}

func (repo *PgRepository) FindMostRecentSeasonByLeague(league *model.League) (*model.Season, error) {
	row := repo.manager.db.QueryRow(`SELECT s.id, s.league_id, s.name, s.start_date, s.end_date, s.created, s.modified
        FROM season s
        WHERE s.league_id = $1
        ORDER BY s.end_date DESC
        LIMIT 1`, league.ID)

	season, err := marshal.Season(row)
	if err != nil {
		return &model.Season{}, err
	}

	return season, err
}
