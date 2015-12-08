package model

import (
	"errors"
	"time"

	"leaguelog/Godeps/_workspace/src/github.com/asaskevich/govalidator"
)

func (league *League) Validate(repo Repository) error {
	if league.Name == "" {
		return errors.New("Cannot create League without name.")
	}

	if league.Sport == "" {
		return errors.New("Cannot create League without sport.")
	}

	return nil
}

func (season *Season) Validate(repo Repository) error {
	if season.League == nil {
		return errors.New("Cannot create Season without League.")
	}

	if season.Name == "" {
		return errors.New("Cannot create Season without Name.")
	}

	t := time.Time{}
	if season.StartDate == t {
		return errors.New("Cannot create Season without start date.")
	}

	if season.EndDate == t {
		return errors.New("Cannot create Season without end date.")
	}

	if season.StartDate.After(season.EndDate) {
		return errors.New("Season start date cannot be after end date.")
	}

	return nil
}

func (team *Team) Validate(repo Repository) error {
	if team.League == nil {
		return errors.New("Cannot create Team without League.")
	}

	if team.Name == "" {
		return errors.New("Cannot create Team without Name.")
	}

	return nil
}

func (game *Game) Validate(repo Repository) error {
	if game.Season == nil {
		return errors.New("Cannot create Game without Season.")
	}

	if game.HomeTeam == nil {
		return errors.New("Cannot create Game without Home team.")
	}

	if game.AwayTeam == nil {
		return errors.New("Cannot create Game without Away team.")
	}

	t := time.Time{}
	if game.StartTime == t {
		return errors.New("Cannot create Game without start time.")
	}

	return nil
}

func (standing *Standing) Validate(repo Repository) error {
	if standing.Season == nil {
		return errors.New("Cannot create Standing without Season.")
	}

	if standing.Team == nil {
		return errors.New("Cannot create Standing without Team.")
	}

	return nil
}

func (user *User) Validate(repo Repository) error {
	if user.Email == "" {
		return errors.New("Cannot create User without email.")
	}

	if !govalidator.IsEmail(user.Email) {
		return UserInvalidEmail
	}

	if user.Password == "" {
		return UserInvalidPassword
	}

	u, _ := repo.FindUserByEmail(user.Email)
	if u != nil {
		return UserDuplicateEmail
	}

	return nil
}
