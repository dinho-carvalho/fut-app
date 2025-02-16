package models

import (
	"time"

	"fut-app/internal/database"
)

type Match struct {
	database.Model
	Date time.Time `gorm:"not null"`
}
