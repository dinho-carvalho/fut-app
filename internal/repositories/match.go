package repositories

import (
	"fut-app/internal/database/models"
	"gorm.io/gorm"
)

type (
	matchRepository struct {
		db *gorm.DB
	}
	MatchRepository interface {
		CreateMatch(models.Match) error
		GetMatches() []models.Match
		GetMatchByID(int) (models.Match, error)
		UpdateMatch(models.Match) error
		DeleteMatch(int) error
	}
)

func NewMatch(db *gorm.DB) MatchRepository {
	return &matchRepository{
		db: db,
	}
}

func (p *matchRepository) CreateMatch(m models.Match) error {
	return p.db.Create(&m).Error
}

func (p *matchRepository) GetMatches() []models.Match {
	var m []models.Match
	p.db.Find(&m)

	return m
}

func (p *matchRepository) GetMatchByID(id int) (models.Match, error) {
	var m models.Match
	err := p.db.First(&m, id).Error

	return m, err
}

func (p *matchRepository) UpdateMatch(m models.Match) error {
	return p.db.Save(&m).Error
}

func (p *matchRepository) DeleteMatch(id int) error {
	return p.db.Delete(&models.Match{}, id).Error
}
