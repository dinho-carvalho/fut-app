package models

import (
	"fut-app/internal/database"
)

type Player struct {
	database.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Stats JSONB  `gorm:"type:jsonb;default:'{}'"`
}
