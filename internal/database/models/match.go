package models

import (
	"time"

	"fut-app/internal/database"
)

type Match struct {
	database.Model
	Date       time.Time `gorm:"not null"`
	Location   string    `gorm:"type:varchar(100)"`
	TeamA      string    `gorm:"type:varchar(100)"`
	TeamB      string    `gorm:"type:varchar(100)"`
	ScoreTeamA int
	ScoreTeamB int
}
