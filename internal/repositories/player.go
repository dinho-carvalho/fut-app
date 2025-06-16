package repositories

import (
	"errors"
	"fmt"
	"log/slog"

	"fut-app/internal/database/models"
	"fut-app/internal/domain"

	"gorm.io/gorm"
)

type (
	playerRepository struct {
		db     *gorm.DB
		logger *slog.Logger
	}
	PlayerRepository interface {
		CreatePlayer(domain.Player) error
		GetPlayers() []models.Player
		GetPlayerByID(uint) (*models.Player, error)
		UpdatePlayer(models.Player) error
		DeletePlayer(uint) error
	}
)

func NewPlayer(DB *gorm.DB, l *slog.Logger) PlayerRepository {
	return &playerRepository{
		db:     DB,
		logger: l,
	}
}

func (p *playerRepository) CreatePlayer(player domain.Player) error {
	positions, err := p.getPositions(player)
	if err != nil {
		return err
	}
	modelPlayer := models.Player{
		Name:     player.Name,
		Stats:    player.Stats,
		Position: positions,
	}
	err = p.db.Create(&modelPlayer).Error
	if err != nil {
		p.logger.Error("erro ao criar jogador",
			slog.String("name", player.Name),
			slog.String("error", err.Error()),
		)
	}
	return err
}

func (p *playerRepository) GetPlayers() []models.Player {
	var players []models.Player
	p.db.Find(&players)

	return players
}

func (p *playerRepository) GetPlayerByID(id uint) (*models.Player, error) {
	var player models.Player
	err := p.db.First(&player, id).Error

	return &player, err
}

func (p *playerRepository) UpdatePlayer(player models.Player) error {
	return p.db.Save(&player).Error
}

func (p *playerRepository) DeletePlayer(id uint) error {
	return p.db.Delete(&models.Player{}, id).Error
}

func (p *playerRepository) getPositions(player domain.Player) ([]models.Position, error) {
	var positions []models.Position
	for _, posName := range player.Position {
		var position models.Position
		if err := p.db.Where("name = ?", posName).First(&position).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				p.logger.Warn("position not founded", slog.String("name", posName))
				return nil, fmt.Errorf("position '%s' not found", posName)
			}
			p.logger.Error("error while fetching position",
				slog.String("name", posName),
				slog.String("error", err.Error()),
			)
			return nil, fmt.Errorf("error while fetching position '%s': %w", posName, err)
		}
		positions = append(positions, position)
	}
	return positions, nil
}
