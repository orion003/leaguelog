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
	Id       int       `json:"id"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

type League struct {
	Model `json:"model"`
	Name  string `json:"name"`
	Sport string `json:"sport"`
}

type Season struct {
	Model      `json:"model"`
	League     *League   `json:"league"`
	Name       string    `json:"name"`
	Start_date time.Time `json:"start_date"`
	End_date   time.Time `json:"end_date"`
}

type Team struct {
	Model  `json:"model"`
	League *League `json:"league"`
	Name   string  `json:"name"`
}

type Game struct {
	Model      `json:"model"`
	Season     *Season   `json:"season"`
	Start_time time.Time `json:"start_time"`
	Home_team  *Team     `json:"home_team"`
	Away_team  *Team     `json:"away_team"`
	Home_score int       `json:"home_score"`
	Away_score int       `json:"away_score"`
}

type User struct {
	Model `json:"model"`
	Email string `json:"email"`
}
