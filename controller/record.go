package controller

import (
    "time"
    
    "leaguelog/model"
)

type Schedule struct {
    StartDate time.Time `json:"start_date"`
    Games []model.Game `json:"games"`
}   