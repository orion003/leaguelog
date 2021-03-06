package postgres

import (
	"time"

	"leaguelog/database/marshal"
	"leaguelog/model"
)

func (repo *PgRepository) CreateGame(game *model.Game) error {
	err := game.Validate(repo)
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO game(season_id, start_time, home_team_id, away_team_id, created, modified)
	    VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		game.Season.ID, game.StartTime, game.HomeTeam.ID, game.AwayTeam.ID, t, t).Scan(&id)

	if err != nil {
		return err
	}

	game.ID = id
	game.Created = t
	game.Modified = t

	return nil
}

func (repo *PgRepository) FindGameByID(id int) (*model.Game, error) {
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

func (repo *PgRepository) FindAllGamesAfterDateBySeason(season *model.Season, after *time.Time) ([]model.Game, error) {
	rows, err := repo.manager.db.Query(`
	    SELECT g.id, g.season_id, g.start_time, g.home_score, g.away_score, g.created, g.modified,
	        t1.id, t1.league_id, t1.name, t1.created, t1.modified,
	        t2.id, t2.league_id, t2.name, t2.created, t2.modified
        FROM game g
        INNER JOIN team t1 on g.home_team_id = t1.id
        INNER JOIN team t2 on g.away_team_id = t2.id
        WHERE g.season_id = $1
            AND g.start_time >= $2
        ORDER BY g.start_time ASC`, season.ID, after)

	if err != nil {
		return []model.Game{}, err
	}

	var games []model.Game
	for rows.Next() {
		game, err := marshal.GameWithTeams(rows)
		if err != nil {
			return []model.Game{}, err
		}

		if games == nil {
			games = make([]model.Game, 0, 1)
		}

		games = append(games, *game)
	}

	err = rows.Err()
	if err != nil {
		return []model.Game{}, err
	}

	return games, nil
}
