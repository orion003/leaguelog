package postgres

import (
	"time"

	"recleague/database/marshal"
	"recleague/model"
)

func (repo *PgGameRepository) Create(game *model.Game) error {
	err := game.Validate()
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO game(season_id, start_time, home_team_id, away_team_id, created, modified) 
	    VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		game.Season.Id, game.Start_time, game.Home_team.Id, game.Away_team.Id, t, t).Scan(&id)

	if err != nil {
		return err
	}

	game.Id = id
	game.Created = t
	game.Modified = t

	return nil
}

func (repo *PgGameRepository) FindById(id int) (*model.Game, error) {
	row := repo.manager.db.QueryRow(`
	    SELECT g.id, g.season_id, g.start_time, g.home_team_id, g.away_team_id, g.home_score, g.away_score, g.created, g.modified
        FROM game g
        WHERE g.id = $1`, id)

	game, err := marshal.Game(row)
	if err != nil {
		return &model.Game{}, err
	}

	return game, nil
}
