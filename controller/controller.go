package controller

import (
	"net/http"
	
	"recleague/model"
)

struct Controller {
    League *model.LeagueRepository
    Season *model.SeasonRepository
    Team *model.TeamRepository
    Game *model.GameRepository
    User *model.UserRepository
}

func NewController(repos map[string]interface) *Controller {
    c := Controller {
        League: repos["league"],
        Season: repos["season"],
        Team: repos["team"],
        Game: repos["game"],
        User: repos["user"],
    }
}

func (c *Controller) AddEmail(w http.ResponseWriter, r *http.Request) {
    
}
