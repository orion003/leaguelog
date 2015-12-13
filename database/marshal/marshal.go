package marshal

import (
	"leaguelog/model"
)

type MultiScanner interface {
	Scan(dest ...interface{}) error
}

func League(row MultiScanner) (*model.League, error) {
	//id, name, sport, created, modified
	var league model.League

	err := row.Scan(&league.ID, &league.Name, &league.Sport, &league.Created, &league.Modified)
	if err != nil {
		return nil, err
	}

	return &league, nil
}

func Season(row MultiScanner) (*model.Season, error) {
	//s.id, s.name, s.start_date, s.end_date, s.created, s.modified, l.id, l.name, l.sport, l.created, l.modified
	var season model.Season
	var league model.League

	err := row.Scan(&season.ID, &league.ID, &season.Name, &season.StartDate, &season.EndDate, &season.Created, &season.Modified)

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

	err := row.Scan(&team.ID, &league.ID, &team.Name, &team.Created, &team.Modified)

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
		&standing.ID, &season.ID, &standing.Wins, &standing.Losses, &standing.Ties, &standing.Created, &standing.Modified,
		&team.ID, &league.ID, &team.Name, &team.Created, &team.Modified)

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

	err := row.Scan(&game.ID, &season.ID, &game.StartTime, &home.ID, &away.ID, &game.HomeScore, &game.AwayScore, &game.Created, &game.Modified)

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

	err := row.Scan(&game.ID, &season.ID, &game.StartTime, &game.HomeScore, &game.AwayScore, &game.Created, &game.Modified,
		&home.ID, &homeLeague.ID, &home.Name, &home.Created, &home.Modified,
		&away.ID, &awayLeague.ID, &away.Name, &away.Created, &away.Modified)

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

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Salt, &user.Created, &user.Modified)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
