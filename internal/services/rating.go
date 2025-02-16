package services

import (
	"fut-app/internal/database/models"
	"fut-app/internal/repositories"
)

type RatingService struct {
	repo repositories.RatingRepository
}

func NewRatingService(repo repositories.RatingRepository) *RatingService {
	return &RatingService{repo: repo}
}

func (s *RatingService) CreateRating(rating models.Rating) error {
	return s.repo.CreateRating(rating)
}

func (s *RatingService) GetAllRatings() []models.Rating {
	return s.repo.GetRatings()
}

func (s *RatingService) GetRatingByID(id int) (models.Rating, error) {
	return s.repo.GetRatingByID(id)
}

func (s *RatingService) UpdateRating(rating models.Rating) error {
	return s.repo.UpdateRating(rating)
}

func (s *RatingService) DeleteRating(id int) error {
	return s.repo.DeleteRating(id)
}
