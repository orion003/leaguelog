package postgres

import (
	"time"

	"recleague/database/marshal"
	"recleague/model"
)

func NewPgStandingRepository(manager *PgManager) *PgStandingRepository {
	repo := &PgStandingRepository{
		manager: manager,
	}

	return repo
}

func (repo *PgStandingRepository) Create(standing *model.Standing) error {
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

func (repo *PgStandingRepository) FindAllBySeason(season *model.Season) ([]model.Standing, error) {
	rows, err := repo.manager.db.Query(`SELECT s.id, s.season_id, s.team_id, s.wins, s.losses, s.ties, s.created, s.modified 
        FROM standing s
        WHERE s.season_id = $1`, season.Id)

	if err != nil {
		return []model.Standing{}, err
	}

	var standings []model.Standing
	for rows.Next() {
		standing, err := marshal.Standing(rows)
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
