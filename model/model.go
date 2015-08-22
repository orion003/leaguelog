package model

import (
	"time"
)

type LeagueRepository interface {
	Create(league *League)
	FindById(id int)
	FindAll()
}

type SeasonRepository interface {
	Create(season *Season)
	FindById(id int)
}

type TeamRepository interface {
	Create(team *Team)
	FindById(id int)
}

type GameRepository interface {
	Create(game *Game)
	FindById(id int)
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
