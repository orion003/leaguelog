package postgres

import (
	"time"

	"leaguelog/database/marshal"
	"leaguelog/model"
)

func (repo *PgRepository) CreateStanding(standing *model.Standing) error {
	err := standing.Validate(repo)
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO standing(season_id, team_id, wins, losses, ties, created, modified)
	    VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		standing.Season.Id, standing.Team.Id, standing.Wins, standing.Losses, standing.Ties, t, t).Scan(&id)

	if err != nil {
		return err
	}

	standing.Id = id
	standing.Created = t
	standing.Modified = t

	return nil
}

func (repo *PgRepository) FindAllStandingsBySeason(season *model.Season) ([]model.Standing, error) {
	rows, err := repo.manager.db.Query(`SELECT s.id, s.season_id, s.wins, s.losses, s.ties, s.created, s.modified,
	        t.id, t.league_id, t.name, t.created, t.modified
        FROM standing s
        INNER JOIN team t ON s.team_id = t.id
        WHERE s.season_id = $1
        ORDER BY s.wins DESC, s.ties DESC, s.losses DESC, t.name ASC`, season.Id)

	if err != nil {
		return []model.Standing{}, err
	}

	var standings []model.Standing
	for rows.Next() {
		standing, err := marshal.StandingWithTeams(rows)
		if err != nil {
			return []model.Standing{}, err
		}

		if standings == nil {
			standings = make([]model.Standing, 0, 1)
		}

		standings = append(standings, *standing)
	}

	err = rows.Err()
	if err != nil {
		return []model.Standing{}, err
	}

	return standings, nil
}
