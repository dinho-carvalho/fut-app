package services

import (
	"errors"

	"fut-app/internal/domain"

	"fut-app/internal/database/models"
	"fut-app/internal/repositories"
)

type (
	PlayerService interface {
		CreatePlayer(domain.Player) error
		GetAllPlayers() []models.Player
		GetPlayerByID(uint) (*models.Player, error)
		UpdatePlayer(models.Player, uint) error
		DeletePlayer(uint) error
	}
	playerimpl struct {
		repo repositories.PlayerRepository
	}
)

func NewPlayerService(repo repositories.PlayerRepository) PlayerService {
	return &playerimpl{repo: repo}
}

func (s *playerimpl) CreatePlayer(player domain.Player) error {
	return s.repo.CreatePlayer(player)
}

func (s *playerimpl) GetAllPlayers() []models.Player {
	return s.repo.GetPlayers()
}

func (s *playerimpl) GetPlayerByID(id uint) (*models.Player, error) {
	return s.repo.GetPlayerByID(id)
}

func (s *playerimpl) UpdatePlayer(player models.Player, id uint) error {
	p, err := s.repo.GetPlayerByID(id)
	if err != nil {
		return errors.New("error") // TODO refatorar
	}
	if p == nil {
		return errors.New("error") // TODO refatorar
	}
	return s.repo.UpdatePlayer(player)
}

func (s *playerimpl) DeletePlayer(id uint) error {
	return s.repo.DeletePlayer(id)
}
