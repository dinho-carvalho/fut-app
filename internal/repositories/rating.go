package repositories

import (
	"fut-app/internal/database/models"
	"gorm.io/gorm"
)

type (
	ratingRepository struct {
		db *gorm.DB
	}
	RatingRepository interface {
		CreateRating(models.Rating) error
		GetRatings() []models.Rating
		GetRatingByID(int) (models.Rating, error)
		UpdateRating(models.Rating) error
		DeleteRating(int) error
	}
)

func NewRating(db *gorm.DB) RatingRepository {
	return &ratingRepository{
		db: db,
	}
}

func (r *ratingRepository) CreateRating(rating models.Rating) error {
	return r.db.Create(&rating).Error
}

func (r *ratingRepository) GetRatings() []models.Rating {
	var rating []models.Rating
	r.db.Find(&rating)

	return rating
}

func (r *ratingRepository) GetRatingByID(id int) (models.Rating, error) {
	var rating models.Rating
	err := r.db.First(&rating, id).Error

	return rating, err
}

func (r *ratingRepository) UpdateRating(rating models.Rating) error {
	return r.db.Save(&rating).Error
}

func (r *ratingRepository) DeleteRating(id int) error {
	return r.db.Delete(&models.Rating{}, id).Error
}
