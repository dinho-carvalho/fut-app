package repositories

import (
	"fut-app/internal/database/models"

	"gorm.io/gorm"
)

type (
	playerRepository struct {
		db *gorm.DB
	}
	PlayerRepository interface {
		CreatePlayer(models.Player) error
		GetPlayers() []models.Player
		GetPlayerByID(int) (*models.Player, error)
		UpdatePlayer(models.Player) error
		DeletePlayer(int) error
	}
)

func NewPlayer(DB *gorm.DB) PlayerRepository {
	return &playerRepository{
		db: DB,
	}
}

func (p *playerRepository) CreatePlayer(player models.Player) error {
	return p.db.Create(&player).Error
}

func (p *playerRepository) GetPlayers() []models.Player {
	var players []models.Player
	p.db.Find(&players)

	return players
}

func (p *playerRepository) GetPlayerByID(id int) (*models.Player, error) {
	var player models.Player
	err := p.db.First(&player, id).Error

	return &player, err
}

func (p *playerRepository) UpdatePlayer(player models.Player) error {
	return p.db.Save(&player).Error
}

func (p *playerRepository) DeletePlayer(id int) error {
	return p.db.Delete(&models.Player{}, id).Error
}
