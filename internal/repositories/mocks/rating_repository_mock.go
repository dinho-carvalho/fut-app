package mocks

import (
	"fut-app/internal/database/models"
)

type RatingRepositoryMock struct {
	CreateRatingFunc  func(models.Rating) error
	GetRatingsFunc    func() []models.Rating
	GetRatingByIDFunc func(int) (models.Rating, error)
	UpdateRatingFunc  func(models.Rating) error
	DeleteRatingFunc  func(int) error
}

func (m *RatingRepositoryMock) CreateRating(rating models.Rating) error {
	if m.CreateRatingFunc != nil {
		return m.CreateRatingFunc(rating)
	}
	return nil
}

func (m *RatingRepositoryMock) GetRatings() []models.Rating {
	if m.GetRatingsFunc != nil {
		return m.GetRatingsFunc()
	}
	return []models.Rating{}
}

func (m *RatingRepositoryMock) GetRatingByID(id int) (models.Rating, error) {
	if m.GetRatingByIDFunc != nil {
		return m.GetRatingByIDFunc(id)
	}
	return models.Rating{}, nil
}

func (m *RatingRepositoryMock) UpdateRating(rating models.Rating) error {
	if m.UpdateRatingFunc != nil {
		return m.UpdateRatingFunc(rating)
	}
	return nil
}

func (m *RatingRepositoryMock) DeleteRating(id int) error {
	if m.DeleteRatingFunc != nil {
		return m.DeleteRatingFunc(id)
	}
	return nil
}
