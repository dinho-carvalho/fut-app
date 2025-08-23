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
		GetMatchByID(uint) (*models.Match, error)
		UpdateMatch(models.Match) error
		DeleteMatch(uint) error
	}
)

func NewMatch(DB *gorm.DB) MatchRepository {
	return &matchRepository{
		db: DB,
	}
}

func (p *matchRepository) CreateMatch(match models.Match) error {
	return p.db.Create(&match).Error
}

func (p *matchRepository) GetMatches() []models.Match {
	var matches []models.Match
	p.db.Find(&matches)

	return matches
}

func (p *matchRepository) GetMatchByID(id uint) (*models.Match, error) {
	var match models.Match
	err := p.db.First(&match, id).Error

	return &match, err
}

func (p *matchRepository) UpdateMatch(match models.Match) error {
	return p.db.Save(&match).Error
}

func (p *matchRepository) DeleteMatch(id uint) error {
	return p.db.Delete(&models.Match{}, id).Error
}
