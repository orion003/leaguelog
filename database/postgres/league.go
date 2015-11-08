package postgres

import (
	"time"

	"leaguelog/database/marshal"
	"leaguelog/model"
)

func (repo *PgRepository) CreateLeague(league *model.League) error {
	err := league.Validate(repo)
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO league(name, sport, created, modified)
	    VALUES($1, $2, $3, $4) RETURNING id`,
		league.Name, league.Sport, t, t).Scan(&id)

	if err != nil {
		return err
	}

	league.Id = id
	league.Created = t
	league.Modified = t

	return nil
}

func (repo *PgRepository) FindAllLeagues() ([]model.League, error) {
	rows, err := repo.manager.db.Query(`SELECT id, name, sport, created, modified
        FROM league`)

	if err != nil {
		return []model.League{}, err
	}

	var leagues []model.League
	for rows.Next() {
		league, err := marshal.League(rows)
		if err != nil {
			return []model.League{}, err
		}

		if leagues == nil {
			leagues = make([]model.League, 0, 1)
		}

		leagues = append(leagues, *league)
	}

	err = rows.Err()
	if err != nil {
		return []model.League{}, err
	}

	return leagues, nil
}

func (repo *PgRepository) FindLeagueById(id int) (*model.League, error) {
	row := repo.manager.db.QueryRow(`SELECT id, name, sport, created, modified
        FROM league
        WHERE id = $1`, id)

	league, err := marshal.League(row)
	if err != nil {
		return nil, err
	}

	return league, nil
}
