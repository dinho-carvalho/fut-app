package models

import (
	"fut-app/internal/database"
)

type Player struct {
	database.Model
	Name     string     `gorm:"type:varchar(100);not null"`
	Position []Position `gorm:"many2many:player_positions;"`
	Stats    *JSONB     `gorm:"type:jsonb;default:'{}'"`
}

type Position struct {
	database.Model
	Name string `gorm:"type:varchar(50);not null;uniqueIndex"`
}
