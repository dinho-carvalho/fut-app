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

	players := repo.GetPlayers()
	if len(players) != 1 || players[0].Name != "Sócrates" {
		t.Errorf("GetPlayers() = %v, want 1 player named Sócrates", players)
	}

	got, err := repo.GetPlayerByID(players[0].ID)
	if err != nil {
		t.Errorf("GetPlayerByID() error = %v", err)
	}
	if got.Name != "Sócrates" {
		t.Errorf("GetPlayerByID() = %v, want Name Sócrates", got)
	}
}

func TestPlayerRepository_UpdatePlayer(t *testing.T) {
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
		Name:     "Rivelino",
		Stats:    stats,
		Position: []string{"Meio-campo"},
	}

	createdPlayer, err := repo.CreatePlayer(player)
	if err != nil {
		t.Fatalf("CreatePlayer() error = %v", err)
	}

	// Update the player
	createdPlayer.Name = "Rivelino Atualizado"
	createdPlayer.Stats["velocidade"] = 100

	err = repo.UpdatePlayer(*createdPlayer)
	if err != nil {
		t.Fatalf("UpdatePlayer() error = %v", err)
	}

	players := repo.GetPlayers()
	if len(players) == 0 {
		t.Fatalf("No players found after update")
	}

	got, err := repo.GetPlayerByID(players[0].ID)
	if err != nil {
		t.Errorf("GetPlayerByID() error = %v", err)
	}
	if got.Name != "Rivelino Atualizado" {
		t.Errorf("UpdatePlayer() = %v, want Name Rivelino Atualizado", got)
	}
}

func TestPlayerRepository_DeletePlayer(t *testing.T) {
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
		Name:     "Garrincha",
		Stats:    stats,
		Position: []string{"Atacante"},
	}

	createdPlayer, err := repo.CreatePlayer(player)
	if err != nil {
		t.Fatalf("CreatePlayer() error = %v", err)
	}

	err = repo.DeletePlayer(createdPlayer.ID)
	if err != nil {
		t.Fatalf("DeletePlayer() error = %v", err)
	}

	_, err = repo.GetPlayerByID(createdPlayer.ID)
	if err == nil {
		t.Errorf("GetPlayerByID() after delete should return error, got nil")
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

func TestPlayerRepository_GetPlayersEmpty(t *testing.T) {
	db := setupTestDB(t)
	logger := slog.Default()
	repo := NewPlayer(db, logger)

	players := repo.GetPlayers()
	if len(players) != 0 {
		t.Errorf("GetPlayers() = %v, want empty slice", players)
	}
}

func TestPlayerRepository_GetPlayerByIDNotFound(t *testing.T) {
	db := setupTestDB(t)
	logger := slog.Default()
	repo := NewPlayer(db, logger)

	_, err := repo.GetPlayerByID(999)
	if err == nil {
		t.Errorf("GetPlayerByID() should return error for non-existent player, got nil")
	}
}
