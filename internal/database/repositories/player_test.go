package repositories

import (
	"log/slog"
	"testing"

	"fut-app/internal/database/models"
	"fut-app/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Enable JSON support for SQLite
	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		t.Fatalf("failed to enable foreign keys: %v", err)
	}

	if err := db.AutoMigrate(&models.Player{}, &models.Position{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func setupTestDBWithPositions(t *testing.T) (*gorm.DB, []models.Position) {
	db := setupTestDB(t)

	// Create test positions
	positions := []models.Position{
		{Name: "Atacante"},
		{Name: "Meio-campo"},
		{Name: "Zagueiro"},
		{Name: "Goleiro"},
	}

	for i := range positions {
		if err := db.Create(&positions[i]).Error; err != nil {
			t.Fatalf("failed to create position: %v", err)
		}
	}

	return db, positions
}

func TestPlayerRepository_CreateAndGetPlayer(t *testing.T) {
	db, _ := setupTestDBWithPositions(t)
	logger := slog.Default()
	repo := NewPlayer(db, logger)

	stats := make(map[string]interface{})
	stats["velocidade"] = 99
	stats["drible"] = 95
	stats["finalizacao"] = 90
	stats["passe"] = 88
	stats["defesa"] = 60
	stats["fisico"] = 85

	player := domain.Player{
		Name:     "Sócrates",
		Stats:    stats,
		Position: []string{"Meio-campo"},
	}

	createdPlayer, err := repo.CreatePlayer(player)
	if err != nil {
		t.Fatalf("CreatePlayer() error = %v", err)
	}

	if createdPlayer.Name != "Sócrates" {
		t.Errorf("CreatePlayer() name = %v, want Sócrates", createdPlayer.Name)
	}

	if len(createdPlayer.Position) != 1 || createdPlayer.Position[0] != "Meio-campo" {
		t.Errorf("CreatePlayer() position = %v, want [Meio-campo]", createdPlayer.Position)
	}
}

func TestPlayerRepository_CreatePlayerWithMultiplePositions(t *testing.T) {
	db, _ := setupTestDBWithPositions(t)
	logger := slog.Default()
	repo := NewPlayer(db, logger)

	stats := make(map[string]interface{})
	stats["velocidade"] = 99
	stats["drible"] = 95
	stats["finalizacao"] = 90
	stats["passe"] = 88
	stats["defesa"] = 60
	stats["fisico"] = 85

	player := domain.Player{
		Name:     "Pelé",
		Stats:    stats,
		Position: []string{"Atacante", "Meio-campo"},
	}

	createdPlayer, err := repo.CreatePlayer(player)
	if err != nil {
		t.Fatalf("CreatePlayer() error = %v", err)
	}

	if len(createdPlayer.Position) != 2 {
		t.Errorf("CreatePlayer() position count = %d, want 2", len(createdPlayer.Position))
	}

	expectedPositions := []string{"Atacante", "Meio-campo"}
	for _, expectedPos := range expectedPositions {
		found := false
		for _, actualPos := range createdPlayer.Position {
			if actualPos == expectedPos {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("CreatePlayer() missing position %s, got %v", expectedPos, createdPlayer.Position)
		}
	}
}

func TestPlayerRepository_CreatePlayerWithInvalidPosition(t *testing.T) {
	db, _ := setupTestDBWithPositions(t)
	logger := slog.Default()
	repo := NewPlayer(db, logger)

	stats := make(map[string]interface{})
	stats["velocidade"] = 99
	stats["drible"] = 95
	stats["finalizacao"] = 90
	stats["passe"] = 88
	stats["defesa"] = 60
	stats["fisico"] = 85

	player := domain.Player{
		Name:     "Test Player",
		Stats:    stats,
		Position: []string{"Posição Inexistente"},
	}

	_, err := repo.CreatePlayer(player)
	if err == nil {
		t.Errorf("CreatePlayer() should return error for invalid position, got nil")
	}

	expectedError := "position 'Posição Inexistente' not found"
	if err.Error() != expectedError {
		t.Errorf("CreatePlayer() error = %v, want %s", err, expectedError)
	}
}
