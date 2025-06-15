package services

import (
	"errors"
	"reflect"
	"testing"

	"fut-app/internal/database/models"
)

// Mock do PlayerRepository
type mockPlayerRepository struct {
	CreatePlayerFn  func(models.Player) error
	GetPlayersFn    func() []models.Player
	GetPlayerByIDFn func(uint) (*models.Player, error)
	UpdatePlayerFn  func(models.Player) error
	DeletePlayerFn  func(uint) error
}

func (m *mockPlayerRepository) CreatePlayer(p models.Player) error {
	return m.CreatePlayerFn(p)
}

func (m *mockPlayerRepository) GetPlayers() []models.Player {
	return m.GetPlayersFn()
}

func (m *mockPlayerRepository) GetPlayerByID(id uint) (*models.Player, error) {
	return m.GetPlayerByIDFn(id)
}

func (m *mockPlayerRepository) UpdatePlayer(p models.Player) error {
	return m.UpdatePlayerFn(p)
}

func (m *mockPlayerRepository) DeletePlayer(id uint) error {
	return m.DeletePlayerFn(id)
}

func TestPlayerService_CreatePlayer(t *testing.T) {
	mockRepo := &mockPlayerRepository{
		CreatePlayerFn: func(p models.Player) error {
			if p.Name == "" {
				return errors.New("invalid")
			}
			return nil
		},
	}
	service := NewPlayerService(mockRepo)

	err := service.CreatePlayer(models.Player{Name: "Zico"})
	if err != nil {
		t.Errorf("CreatePlayer() unexpected error: %v", err)
	}

	err = service.CreatePlayer(models.Player{Name: ""})
	if err == nil {
		t.Errorf("CreatePlayer() expected error for empty name")
	}
}

func TestPlayerService_GetAllPlayers(t *testing.T) {
	expected := []models.Player{{Name: "A"}, {Name: "B"}}
	mockRepo := &mockPlayerRepository{
		GetPlayersFn: func() []models.Player {
			return expected
		},
	}
	service := NewPlayerService(mockRepo)

	got := service.GetAllPlayers()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("GetAllPlayers() = %v, want %v", got, expected)
	}
}

func TestPlayerService_GetPlayerByID(t *testing.T) {
	player := &models.Player{Name: "Pelé"}
	mockRepo := &mockPlayerRepository{
		GetPlayerByIDFn: func(id uint) (*models.Player, error) {
			if id == 1 {
				return player, nil
			}
			return nil, errors.New("not found")
		},
	}
	service := NewPlayerService(mockRepo)

	got, err := service.GetPlayerByID(1)
	if err != nil || got != player {
		t.Errorf("GetPlayerByID(1) = %v, %v; want %v, nil", got, err, player)
	}

	got, err = service.GetPlayerByID(2)
	if err == nil || got != nil {
		t.Errorf("GetPlayerByID(2) expected error, got %v, %v", got, err)
	}
}

func TestPlayerService_UpdatePlayer(t *testing.T) {
	player := models.Player{Name: "Romário"}
	mockRepo := &mockPlayerRepository{
		GetPlayerByIDFn: func(id uint) (*models.Player, error) {
			if id == 1 {
				return &player, nil
			}
			if id == 2 {
				return nil, nil
			}
			return nil, errors.New("db error")
		},
		UpdatePlayerFn: func(p models.Player) error {
			if p.Name == "" {
				return errors.New("invalid")
			}
			return nil
		},
	}
	service := NewPlayerService(mockRepo)

	// Caso de sucesso
	err := service.UpdatePlayer(player, 1)
	if err != nil {
		t.Errorf("UpdatePlayer() unexpected error: %v", err)
	}

	// Caso: jogador não encontrado (nil)
	err = service.UpdatePlayer(player, 2)
	if err == nil {
		t.Errorf("UpdatePlayer() expected error for nil player")
	}

	// Caso: erro ao buscar jogador
	err = service.UpdatePlayer(player, 3)
	if err == nil {
		t.Errorf("UpdatePlayer() expected error for db error")
	}
}

func TestPlayerService_DeletePlayer(t *testing.T) {
	mockRepo := &mockPlayerRepository{
		DeletePlayerFn: func(id uint) error {
			if id == 1 {
				return nil
			}
			return errors.New("not found")
		},
	}
	service := NewPlayerService(mockRepo)

	err := service.DeletePlayer(1)
	if err != nil {
		t.Errorf("DeletePlayer(1) unexpected error: %v", err)
	}

	err = service.DeletePlayer(2)
	if err == nil {
		t.Errorf("DeletePlayer(2) expected error")
	}
}
