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

func (p *matchRepository) CreateMatch(player models.Match) error {
	return p.db.Create(&player).Error
}

func (p *matchRepository) GetMatches() []models.Match {
	var players []models.Match
	p.db.Find(&players)

	return players
}

func (p *matchRepository) GetMatchByID(id int) (models.Match, error) {
	var player models.Match
	err := p.db.First(&player, id).Error

	return player, err
}

func (p *matchRepository) UpdateMatch(player models.Match) error {
	return p.db.Save(&player).Error
}

func (p *matchRepository) DeleteMatch(id int) error {
	return p.db.Delete(&models.Match{}, id).Error
}
