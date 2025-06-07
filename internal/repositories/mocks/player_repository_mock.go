package mocks

import (
	"fut-app/internal/database/models"
)

type PlayerRepositoryMock struct {
	CreatePlayerFunc  func(models.Player) error
	GetPlayersFunc    func() []models.Player
	GetPlayerByIDFunc func(int) (*models.Player, error)
	UpdatePlayerFunc  func(models.Player) error
	DeletePlayerFunc  func(int) error
}

func (m *PlayerRepositoryMock) CreatePlayer(player models.Player) error {
	if m.CreatePlayerFunc != nil {
		return m.CreatePlayerFunc(player)
	}
	return nil
}

func (m *PlayerRepositoryMock) GetPlayers() []models.Player {
	if m.GetPlayersFunc != nil {
		return m.GetPlayersFunc()
	}
	return []models.Player{}
}

func (m *PlayerRepositoryMock) GetPlayerByID(id int) (*models.Player, error) {
	if m.GetPlayerByIDFunc != nil {
		return m.GetPlayerByIDFunc(id)
	}
	return nil, nil
}

func (m *PlayerRepositoryMock) UpdatePlayer(player models.Player) error {
	if m.UpdatePlayerFunc != nil {
		return m.UpdatePlayerFunc(player)
	}
	return nil
}

func (m *PlayerRepositoryMock) DeletePlayer(id int) error {
	if m.DeletePlayerFunc != nil {
		return m.DeletePlayerFunc(id)
	}
	return nil
}
