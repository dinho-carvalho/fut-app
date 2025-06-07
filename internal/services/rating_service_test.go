package services

import (
	"errors"
	"testing"

	"fut-app/internal/database/models"
	"fut-app/internal/repositories/mocks"
)

// TestRatingService_CreateRating tests CreateRating with a nil repository.
func TestRatingService_CreateRating(t *testing.T) {
	tests := []struct {
		name    string
		rating  models.Rating
		mockFn  func(mock *mocks.RatingRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			rating: models.Rating{
				MatchID:       1,
				PlayerID:      1,
				RatedPlayerID: 2,
				Finishing:     80,
			},
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.CreateRatingFunc = func(rating models.Rating) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "database error",
			rating: models.Rating{
				MatchID:       1,
				PlayerID:      1,
				RatedPlayerID: 2,
				Finishing:     80,
			},
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.CreateRatingFunc = func(rating models.Rating) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.RatingRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewRatingService(mockRepo)

			err := service.CreateRating(tt.rating)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRating() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestRatingService_GetAllRatings tests GetAllRatings with a nil repository.
func TestRatingService_GetAllRatings(t *testing.T) {
	tests := []struct {
		name     string
		mockFn   func(mock *mocks.RatingRepositoryMock)
		expected []models.Rating
	}{
		{
			name: "success with ratings",
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.GetRatingsFunc = func() []models.Rating {
					return []models.Rating{{MatchID: 1, PlayerID: 1}}
				}
			},
			expected: []models.Rating{{MatchID: 1, PlayerID: 1}},
		},
		{
			name: "empty list",
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.GetRatingsFunc = func() []models.Rating {
					return []models.Rating{}
				}
			},
			expected: []models.Rating{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.RatingRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewRatingService(mockRepo)

			got := service.GetAllRatings()
			if len(got) != len(tt.expected) {
				t.Errorf("GetAllRatings() got = %v ratings, want %v", len(got), len(tt.expected))
			}
		})
	}
}

// TestRatingService_GetRatingByID tests GetRatingByID with a nil repository.
func TestRatingService_GetRatingByID(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		mockFn  func(mock *mocks.RatingRepositoryMock)
		want    models.Rating
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.GetRatingByIDFunc = func(id int) (models.Rating, error) {
					return models.Rating{MatchID: 1, PlayerID: 1}, nil
				}
			},
			want:    models.Rating{MatchID: 1, PlayerID: 1},
			wantErr: false,
		},
		{
			name: "not found",
			id:   999,
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.GetRatingByIDFunc = func(id int) (models.Rating, error) {
					return models.Rating{}, errors.New("not found")
				}
			},
			want:    models.Rating{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.RatingRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewRatingService(mockRepo)

			got, err := service.GetRatingByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRatingByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.MatchID != tt.want.MatchID {
				t.Errorf("GetRatingByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestRatingService_UpdateRating tests UpdateRating with a nil repository.
func TestRatingService_UpdateRating(t *testing.T) {
	tests := []struct {
		name    string
		rating  models.Rating
		mockFn  func(mock *mocks.RatingRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			rating: models.Rating{
				MatchID:       1,
				PlayerID:      1,
				RatedPlayerID: 2,
				Finishing:     90,
			},
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.UpdateRatingFunc = func(rating models.Rating) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "not found",
			rating: models.Rating{
				MatchID:       1,
				PlayerID:      1,
				RatedPlayerID: 2,
				Finishing:     90,
			},
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.UpdateRatingFunc = func(rating models.Rating) error {
					return errors.New("not found")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.RatingRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewRatingService(mockRepo)

			err := service.UpdateRating(tt.rating)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRating() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestRatingService_DeleteRating tests DeleteRating with a nil repository.
func TestRatingService_DeleteRating(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		mockFn  func(mock *mocks.RatingRepositoryMock)
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.DeleteRatingFunc = func(id int) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "not found",
			id:   999,
			mockFn: func(mock *mocks.RatingRepositoryMock) {
				mock.DeleteRatingFunc = func(id int) error {
					return errors.New("not found")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.RatingRepositoryMock{}
			tt.mockFn(mockRepo)
			service := NewRatingService(mockRepo)

			err := service.DeleteRating(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRating() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
