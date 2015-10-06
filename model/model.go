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
	FindMostRecentByLeague(league *League) (*Season, error)
}

type TeamRepository interface {
	Create(team *Team) error
	FindById(id int) (*Team, error)
}

type StandingRepository interface {
	Create(standing *Standing) error
	FindAllBySeason(season *Season) ([]Standing, error)
}

type GameRepository interface {
	Create(game *Game) error
	FindById(id int) (*Game, error)
	FindUpcomingBySeason(season *Season) ([]Game, error)
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
	Model     `json:"model"`
	League    *League   `json:"league"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Team struct {
	Model  `json:"model"`
	League *League `json:"league"`
	Name   string  `json:"name"`
}

type Standing struct {
	Model  `json:"model"`
	Season *Season `json:"season"`
	Team   *Team   `json:"team"`
	Wins   int     `json:"wins"`
	Losses int     `json:"losses"`
	Ties   int     `json:"ties"`
}

type Game struct {
	Model     `json:"model"`
	Season    *Season   `json:"season"`
	StartTime time.Time `json:"start_time"`
	HomeTeam  *Team     `json:"home_team"`
	AwayTeam  *Team     `json:"away_team"`
	HomeScore int       `json:"home_score"`
	AwayScore int       `json:"away_score"`
}

type User struct {
	Model `json:"model"`
	Email string `json:"email"`
}
