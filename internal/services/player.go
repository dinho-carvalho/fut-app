package services

import (
	"errors"

	"fut-app/internal/database/models"
	"fut-app/internal/repositories"
)

type PlayerService struct {
	repo repositories.PlayerRepository
}

func NewPlayerService(repo repositories.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) CreatePlayer(player models.Player) error {
	return s.repo.CreatePlayer(player)
}

func (s *PlayerService) GetAllPlayers() []models.Player {
	return s.repo.GetPlayers()
}

func (s *PlayerService) GetPlayerByID(id int) (*models.Player, error) {
	return s.repo.GetPlayerByID(id)
}

func (s *PlayerService) UpdatePlayer(player models.Player, id int) error {
	p, err := s.repo.GetPlayerByID(id)
	if err != nil {
		return errors.New("error") // TODO refatorar
	}
	if p == nil {
		return errors.New("error") // TODO refatorar
	}
	return s.repo.UpdatePlayer(player)
}

func (s *PlayerService) DeletePlayer(id int) error {
	return s.repo.DeletePlayer(id)
}
