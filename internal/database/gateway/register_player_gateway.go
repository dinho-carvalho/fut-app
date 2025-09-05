package gateway

import (
	"fut-app/internal/database/repositories"
	"fut-app/internal/domain"
	"fut-app/internal/usecase"
)

type (
	registerPlayerGateway struct {
		repo repositories.Player
	}
)

func NewRegisterPlayerGateway(repo repositories.Player) usecase.RegisterPlayerGateway {
	return &registerPlayerGateway{repo: repo}
}

func (g *registerPlayerGateway) Register(player domain.Player) (*domain.Player, error) {
	return g.repo.CreatePlayer(player)
}
