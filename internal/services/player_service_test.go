package services

import (
	"errors"
	"testing"

	"fut-app/internal/database/models"
	"fut-app/internal/repositories/mocks"
)

// TestPlayerService_CreatePlayer tests CreatePlayer with a nil repository.
func TestPlayerService_CreatePlayer(t *testing.T) {
	tests := []struct {
		name    string
		player  models.Player
		mockFn  func(mock *mocks.PlayerRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			player: models.Player{
				Name: "Test Player",
			},
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.CreatePlayerFunc = func(player models.Player) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "error",
			player: models.Player{
				Name: "Test Player",
			},
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.CreatePlayerFunc = func(player models.Player) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.PlayerRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewPlayerService(mockRepo)

			err := service.CreatePlayer(tt.player)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestPlayerService_GetAllPlayers tests GetAllPlayers with a nil repository.
func TestPlayerService_GetAllPlayers(t *testing.T) {
	tests := []struct {
		name     string
		mockFn   func(mock *mocks.PlayerRepositoryMock)
		expected []models.Player
	}{
		{
			name: "success with players",
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.GetPlayersFunc = func() []models.Player {
					return []models.Player{{Name: "Player 1"}, {Name: "Player 2"}}
				}
			},
			expected: []models.Player{{Name: "Player 1"}, {Name: "Player 2"}},
		},
		{
			name: "success empty",
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.GetPlayersFunc = func() []models.Player {
					return []models.Player{}
				}
			},
			expected: []models.Player{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.PlayerRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewPlayerService(mockRepo)

			got := service.GetAllPlayers()
			if len(got) != len(tt.expected) {
				t.Errorf("GetAllPlayers() got = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestPlayerService_GetPlayerByID tests GetPlayerByID with a nil repository.
func TestPlayerService_GetPlayerByID(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		mockFn  func(mock *mocks.PlayerRepositoryMock)
		want    *models.Player
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.GetPlayerByIDFunc = func(id int) (*models.Player, error) {
					return &models.Player{Name: "Player 1"}, nil
				}
			},
			want:    &models.Player{Name: "Player 1"},
			wantErr: false,
		},
		{
			name: "not found",
			id:   1,
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.GetPlayerByIDFunc = func(id int) (*models.Player, error) {
					return nil, errors.New("not found")
				}
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.PlayerRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewPlayerService(mockRepo)

			got, err := service.GetPlayerByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got.Name != tt.want.Name {
				t.Errorf("GetPlayerByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestPlayerService_UpdatePlayer tests UpdatePlayer with a nil repository.
// This should panic when calling GetPlayerByID on the nil repo.
func TestPlayerService_UpdatePlayer(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		player  models.Player
		mockFn  func(mock *mocks.PlayerRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			player: models.Player{
				Name: "Updated Player",
			},
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.GetPlayerByIDFunc = func(id int) (*models.Player, error) {
					return &models.Player{Name: "Original Player"}, nil
				}
				mock.UpdatePlayerFunc = func(player models.Player) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "player not found",
			id:   1,
			player: models.Player{
				Name: "Updated Player",
			},
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.GetPlayerByIDFunc = func(id int) (*models.Player, error) {
					return nil, errors.New("not found")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.PlayerRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewPlayerService(mockRepo)

			err := service.UpdatePlayer(tt.player)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestPlayerService_DeletePlayer tests DeletePlayer with a nil repository.
func TestPlayerService_DeletePlayer(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		mockFn  func(mock *mocks.PlayerRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.DeletePlayerFunc = func(id int) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "error",
			id:   1,
			mockFn: func(mock *mocks.PlayerRepositoryMock) {
				mock.DeletePlayerFunc = func(id int) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.PlayerRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewPlayerService(mockRepo)

			err := service.DeletePlayer(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
