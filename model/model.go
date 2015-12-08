package model

import (
	"time"
)

type Repository interface {
	CreateLeague(league *League) error
	FindLeagueByID(id int) (*League, error)
	FindAllLeagues() ([]League, error)

	CreateSeason(season *Season) error
	FindSeasonByID(id int) (*Season, error)
	FindMostRecentSeasonByLeague(league *League) (*Season, error)

	CreateTeam(team *Team) error
	FindTeamByID(id int) (*Team, error)

	CreateStanding(standing *Standing) error
	FindAllStandingsBySeason(season *Season) ([]Standing, error)

	CreateGame(game *Game) error
	FindGameByID(id int) (*Game, error)
	FindAllGamesAfterDateBySeason(season *Season, t *time.Time) ([]Game, error)

	CreateUser(user *User) error
	FindUserByID(id int) (*User, error)
	FindUserByEmail(email string) (*User, error)
	FindAllUsers() ([]User, error)
}

type Model struct {
	ID       int       `json:"id"`
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
	Model    `json:"model"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
