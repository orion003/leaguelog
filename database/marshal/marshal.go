package marshal

import (
	"recleague/model"
)

type MultiScanner interface {
	Scan(dest ...interface{}) error
}

func League(row MultiScanner) (*model.League, error) {
	//id, name, sport, created, modified
	var league model.League

	err := row.Scan(&league.Id, &league.Name, &league.Sport, &league.Created, &league.Modified)
	if err != nil {
		return &model.League{}, err
	}

	return &league, nil
}

func Season(row MultiScanner) (*model.Season, error) {
	//s.id, s.name, s.start_date, s.end_date, s.created, s.modified, l.id, l.name, l.sport, l.created, l.modified
	var season model.Season
	var league model.League

	err := row.Scan(&season.Id, &season.Name, &season.Start_date, &season.End_date, &season.Created, &season.Modified,
		&league.Id, &league.Name, &league.Sport, &league.Created, &league.Modified)

	if err != nil {
		return &model.Season{}, err
	}

	season.League = &league

	return &season, nil
}

func Team(row MultiScanner) (*model.Team, error) {
	//t.id, t.name, t.created, t.modified, l.id, l.name, l.sport, l.created, l.modified
	var team model.Team
	var league model.League

	err := row.Scan(&team.Id, &team.Name, &team.Created, &team.Modified,
		&league.Id, &league.Name, &league.Sport, &league.Created, &league.Modified)

	if err != nil {
		return &model.Team{}, err
	}

	team.League = &league

	return &team, nil
}

func Game(row MultiScanner) (*model.Game, error) {
	// g.id, g.start_time, g.home_score, g.away_score, g.created, g.modified
	// s.id, s.name, s.start_date, s.end_date, s.created, s.modified,
	// t1.id, t1.name, t1.created, t1.modified,
	// t2.id, t.name, t2.created, t2.modified,
	// l.id, l.name, l.sport, l.created, l.modified
	var game model.Game
	var season model.Season
	var home model.Team
	var away model.Team
	var league model.League

	err := row.Scan(&game.Id, &game.Start_time, &game.Home_score, &game.Away_score, &game.Created, &game.Modified,
		&season.Id, &season.Name, &season.Start_date, &season.End_date, &season.Created, &season.Modified,
		&home.Id, &home.Name, &home.Created, &home.Modified,
		&away.Id, &away.Name, &away.Created, &away.Modified,
		&league.Id, &league.Name, &league.Sport, &league.Created, &league.Modified)

	if err != nil {
		return &model.Game{}, err
	}

	game.Season = &season
	game.Season.League = &league
	game.Home_team = &home
	game.Home_team.League = &league
	game.Away_team = &away
	game.Away_team.League = &league

	return &game, nil
}
