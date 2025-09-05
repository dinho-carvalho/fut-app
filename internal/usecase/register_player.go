package usecase

import (
	"fut-app/internal/domain"
)

type (
	RegisterPlayerUseCase interface {
		Execute(domain.Player) (*domain.Player, error)
	}
	RegisterPlayerGateway interface {
		Register(domain.Player) (*domain.Player, error)
	}
	player struct {
		gateway RegisterPlayerGateway
	}
)

func NewPlayerUseCase(gateway RegisterPlayerGateway) RegisterPlayerUseCase {
	return &player{gateway: gateway}
}

func (uc *player) Execute(player domain.Player) (*domain.Player, error) {
	if err := player.Validate(); err != nil {
		return nil, err
	}
	return uc.gateway.Register(player)
}
