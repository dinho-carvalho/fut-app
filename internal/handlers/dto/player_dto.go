package dto

import (
	"fut-app/internal/domain"
)

type (
	PlayerDTO struct {
		Name     string                 `json:"name" validate:"required"`
		Stats    map[string]interface{} `json:"stats" validate:"required,statslen"`
		Position []string               `json:"positions" validate:"required,min=1,dive,required"`
	}

	PositionDTO struct {
		Name string `json:"name" validate:"required"`
	}
)

func (p *PlayerDTO) ToDomain() domain.Player {
	return domain.Player{
		Name:     p.Name,
		Stats:    p.Stats,
		Position: p.Position,
	}
}
