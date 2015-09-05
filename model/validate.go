package model

import (
	"errors"
	"time"
)

type Validator interface {
	Validate() error
}

func (league *League) Validate() error {
	if league.Name == "" {
		return errors.New("Cannot create League without name.")
	}

	if league.Sport == "" {
		return errors.New("Cannot create League without sport.")
	}

	return nil
}

func (season *Season) Validate() error {
	if season.League == nil {
		return errors.New("Cannot create Season without League.")
	}

	if season.Name == "" {
		return errors.New("Cannot create Season without Name.")
	}

	t := time.Time{}
	if season.Start_date == t {
		return errors.New("Cannot create Season without start date.")
	}

	if season.End_date == t {
		return errors.New("Cannot create Season without end date.")
	}

	if season.Start_date.After(season.End_date) {
		return errors.New("Season start date cannot be after end date.")
	}

	return nil
}

func (team *Team) Validate() error {
	if team.League == nil {
		return errors.New("Cannot create Team without League.")
	}

	if team.Name == "" {
		return errors.New("Cannot create Team without Name.")
	}

	return nil
}

func (game *Game) Validate() error {
	if game.Season == nil {
		return errors.New("Cannot create Game without Season.")
	}

	if game.Home_team == nil {
		return errors.New("Cannot create Game without Home team.")
	}

	if game.Away_team == nil {
		return errors.New("Cannot create Game without Away team.")
	}

	t := time.Time{}
	if game.Start_time == t {
		return errors.New("Cannot create Game without start time.")
	}

	return nil
}

func (user *User) Validate() error {
    if user.Email == "" {
        return errors.New("Cannot create User without email.")   
    }
    
    return nil
}
