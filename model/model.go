package model

import (
	"time"
)

type LeagueRepository interface {
	Create(league *League) error
	FindById(id int) (*League, error)
	FindAll() ([]League, error)
}

type SeasonRepository interface {
	Create(season *Season) error
	FindById(id int) (*Season, error)
}

type TeamRepository interface {
	Create(team *Team) error
	FindById(id int) (*Team, error)
}

type GameRepository interface {
	Create(game *Game) error
	FindById(id int) (*Game, error)
}

type UserRepository interface {
	Create(user *User) error
	FindById(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll() ([]User, error)
}

type Model struct {
	Id       int
	Created  time.Time
	Modified time.Time
}

type League struct {
	Model
	Name  string
	Sport string
}

type Season struct {
	Model
	League     *League
	Name       string
	Start_date time.Time
	End_date   time.Time
}

type Team struct {
	Model
	League *League
	Name   string
}

type Game struct {
	Model
	Season     *Season
	Start_time time.Time
	Home_team  *Team
	Away_team  *Team
	Home_score int
	Away_score int
}

type User struct {
	Model
	Email string
}
