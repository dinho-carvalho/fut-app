package services

import (
	"errors"
	"testing"
	"time"

	"fut-app/internal/database/models"
	"fut-app/internal/repositories/mocks"
	// "fut-app/internal/repositories" // Not strictly needed for type compilation with nil repo
)

// TestMatchService_CreateMatch tests CreateMatch with a nil repository.
func TestMatchService_CreateMatch(t *testing.T) {
	tests := []struct {
		name    string
		match   models.Match
		mockFn  func(mock *mocks.MatchRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			match: models.Match{
				Date: time.Now(),
			},
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.CreateMatchFunc = func(match models.Match) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "error",
			match: models.Match{
				Date: time.Now(),
			},
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.CreateMatchFunc = func(match models.Match) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MatchRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewMatchService(mockRepo)

			err := service.CreateMatch(tt.match)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestMatchService_GetAllMatches tests GetAllMatches with a nil repository.
func TestMatchService_GetAllMatches(t *testing.T) {
	tests := []struct {
		name     string
		mockFn   func(mock *mocks.MatchRepositoryMock)
		expected []models.Match
	}{
		{
			name: "success with matches",
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.GetMatchesFunc = func() []models.Match {
					return []models.Match{{Date: time.Now()}}
				}
			},
			expected: []models.Match{{Date: time.Now()}},
		},
		{
			name: "success empty",
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.GetMatchesFunc = func() []models.Match {
					return []models.Match{}
				}
			},
			expected: []models.Match{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MatchRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewMatchService(mockRepo)

			got := service.GetAllMatches()
			if len(got) != len(tt.expected) {
				t.Errorf("GetAllMatches() got = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchService_GetMatchByID tests GetMatchByID with a nil repository.
func TestMatchService_GetMatchByID(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		mockFn  func(mock *mocks.MatchRepositoryMock)
		want    models.Match
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.GetMatchByIDFunc = func(id int) (models.Match, error) {
					return models.Match{Date: time.Now()}, nil
				}
			},
			want:    models.Match{Date: time.Now()},
			wantErr: false,
		},
		{
			name: "not found",
			id:   1,
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.GetMatchByIDFunc = func(id int) (models.Match, error) {
					return models.Match{}, errors.New("not found")
				}
			},
			want:    models.Match{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MatchRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewMatchService(mockRepo)

			got, err := service.GetMatchByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Date.IsZero() != tt.want.Date.IsZero() {
				t.Errorf("GetMatchByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMatchService_UpdateMatch tests UpdateMatch with a nil repository.
func TestMatchService_UpdateMatch(t *testing.T) {
	tests := []struct {
		name    string
		match   models.Match
		mockFn  func(mock *mocks.MatchRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			match: models.Match{
				Date: time.Now(),
			},
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.UpdateMatchFunc = func(match models.Match) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "error",
			match: models.Match{
				Date: time.Now(),
			},
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.UpdateMatchFunc = func(match models.Match) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MatchRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewMatchService(mockRepo)

			err := service.UpdateMatch(tt.match)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateMatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestMatchService_DeleteMatch tests DeleteMatch with a nil repository.
func TestMatchService_DeleteMatch(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		mockFn  func(mock *mocks.MatchRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.DeleteMatchFunc = func(id int) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "error",
			id:   1,
			mockFn: func(mock *mocks.MatchRepositoryMock) {
				mock.DeleteMatchFunc = func(id int) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MatchRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewMatchService(mockRepo)

			err := service.DeleteMatch(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteMatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
