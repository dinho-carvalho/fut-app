package models

import "fut-app/internal/database"

type Rating struct {
	database.Model
	MatchID       uint `gorm:"not null;index"`
	PlayerID      uint `gorm:"not null;index"`
	RatedPlayerID uint `gorm:"not null;index"`
	Finishing     int  `gorm:"check:finishing BETWEEN 45 AND 99"`
	Passing       int  `gorm:"check:passing BETWEEN 45 AND 99"`
	Speed         int  `gorm:"check:speed BETWEEN 45 AND 99"`
	Defense       int  `gorm:"check:defense BETWEEN 45 AND 99"`
	Stamina       int  `gorm:"check:stamina BETWEEN 45 AND 99"`
	Highlight     int  `gorm:"check:highlight BETWEEN 45 AND 99"`
}
