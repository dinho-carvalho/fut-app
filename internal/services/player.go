package services

import (
	"fmt"

	"fut-app/internal/database/models"
	"fut-app/internal/repositories"
)

type IPlayerService interface {
	CreatePlayer(models.Player) error
	GetAllPlayers() []models.Player
	GetPlayerByID(int) (models.Player, error)
	UpdatePlayer(models.Player) error
	DeletePlayer(int) error
}

type PlayerService struct {
	repo repositories.PlayerRepository
}

func NewPlayerService(repo repositories.PlayerRepository) IPlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) CreatePlayer(player models.Player) error {
	if player.Name == "" {
		return fmt.Errorf("nome do jogador é obrigatório")
	}
	return s.repo.CreatePlayer(player)
}

func (s *PlayerService) GetAllPlayers() []models.Player {
	return s.repo.GetPlayers()
}

func (s *PlayerService) GetPlayerByID(id int) (models.Player, error) {
	player, err := s.repo.GetPlayerByID(id)
	if err != nil {
		return models.Player{}, fmt.Errorf("erro ao buscar jogador: %w", err)
	}
	if player == nil {
		return models.Player{}, fmt.Errorf("jogador não encontrado")
	}
	return *player, nil
}

func (s *PlayerService) UpdatePlayer(player models.Player) error {
	if player.Name == "" {
		return fmt.Errorf("nome do jogador é obrigatório")
	}

	existingPlayer, err := s.repo.GetPlayerByID(int(player.ID))
	if err != nil {
		return fmt.Errorf("erro ao buscar jogador: %w", err)
	}
	if existingPlayer == nil {
		return fmt.Errorf("jogador não encontrado")
	}

	return s.repo.UpdatePlayer(player)
}

func (s *PlayerService) DeletePlayer(id int) error {
	existingPlayer, err := s.repo.GetPlayerByID(id)
	if err != nil {
		return fmt.Errorf("erro ao buscar jogador: %w", err)
	}
	if existingPlayer == nil {
		return fmt.Errorf("jogador não encontrado")
	}

	return s.repo.DeletePlayer(id)
}
