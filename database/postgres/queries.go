package postgres

import (
	"time"

	"recleague/database/marshal"
	"recleague/model"
)

func (repo *PgLeagueRepository) Create(league *model.League) error {
	err := league.Validate()
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

func (repo *PgLeagueRepository) FindAll() ([]model.League, error) {
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
			leagues = make([]model.League, 1, 10)
		}

		leagues = append(leagues, *league)
	}

	err = rows.Err()
	if err != nil {
		return []model.League{}, err
	}

	return leagues, nil
}

func (repo *PgLeagueRepository) FindById(id int) (*model.League, error) {
	row := repo.manager.db.QueryRow(`SELECT id, name, sport, created, modified
        FROM league
        WHERE id = $1`, id)

	league, err := marshal.League(row)
	if err != nil {
		return &model.League{}, err
	}

	return league, nil
}

func (repo *PgSeasonRepository) Create(season *model.Season) error {
	err := season.Validate()
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO season(league_id, name, start_date, end_date, created, modified) 
	    VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		season.League.Id, season.Name, season.Start_date, season.End_date, t, t).Scan(&id)

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

func (repo *PgTeamRepository) Create(team *model.Team) error {
	err := team.Validate()
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
