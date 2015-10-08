package marshal

import (
	"leaguelog.ca/model"
)

type MultiScanner interface {
	Scan(dest ...interface{}) error
}

func League(row MultiScanner) (*model.League, error) {
	//id, name, sport, created, modified
	var league model.League

	err := row.Scan(&league.Id, &league.Name, &league.Sport, &league.Created, &league.Modified)
	if err != nil {
		return nil, err
	}

	return &league, nil
}

func Season(row MultiScanner) (*model.Season, error) {
	//s.id, s.name, s.start_date, s.end_date, s.created, s.modified, l.id, l.name, l.sport, l.created, l.modified
	var season model.Season
	var league model.League

	err := row.Scan(&season.Id, &league.Id, &season.Name, &season.StartDate, &season.EndDate, &season.Created, &season.Modified)

	if err != nil {
		return nil, err
	}

	season.League = &league

	return &season, nil
}

func Team(row MultiScanner) (*model.Team, error) {
	//t.id, t.name, t.created, t.modified, l.id, l.name, l.sport, l.created, l.modified
	var team model.Team
	var league model.League

	err := row.Scan(&team.Id, &league.Id, &team.Name, &team.Created, &team.Modified)

	if err != nil {
		return nil, err
	}

	team.League = &league

	return &team, nil
}

func StandingWithTeams(row MultiScanner) (*model.Standing, error) {
	var standing model.Standing
	var season model.Season
	var team model.Team
	var league model.League

	err := row.Scan(
		&standing.Id, &season.Id, &standing.Wins, &standing.Losses, &standing.Ties, &standing.Created, &standing.Modified,
		&team.Id, &league.Id, &team.Name, &team.Created, &team.Modified)

	if err != nil {
		return nil, err
	}

	standing.Season = &season
	standing.Team = &team
	standing.Team.League = &league

	return &standing, nil
}

func Game(row MultiScanner) (*model.Game, error) {
	var game model.Game
	var season model.Season
	var home model.Team
	var away model.Team

	err := row.Scan(&game.Id, &season.Id, &game.StartTime, &home.Id, &away.Id, &game.HomeScore, &game.AwayScore, &game.Created, &game.Modified)

	if err != nil {
		return nil, err
	}

	game.Season = &season
	game.HomeTeam = &home
	game.AwayTeam = &away

	return &game, nil
}

func GameWithTeams(row MultiScanner) (*model.Game, error) {
	var game model.Game
	var homeLeague model.League
	var awayLeague model.League
	var season model.Season
	var home model.Team
	var away model.Team

	err := row.Scan(&game.Id, &season.Id, &game.StartTime, &game.HomeScore, &game.AwayScore, &game.Created, &game.Modified,
		&home.Id, &homeLeague.Id, &home.Name, &home.Created, &home.Modified,
		&away.Id, &awayLeague.Id, &away.Name, &away.Created, &away.Modified)

	if err != nil {
		return nil, err
	}

	home.League = &homeLeague
	away.League = &awayLeague

	game.Season = &season
	game.HomeTeam = &home
	game.AwayTeam = &away

	return &game, nil
}

func User(row MultiScanner) (*model.User, error) {
	var user model.User

	err := row.Scan(&user.Id, &user.Email, &user.Created, &user.Modified)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
