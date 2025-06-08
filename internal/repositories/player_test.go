package repositories

import (
	"testing"

	"fut-app/internal/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	if err := db.AutoMigrate(&models.Player{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestPlayerRepository_CreateAndGetPlayer(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPlayer(db)

	stats := make(map[string]int)
	stats["velocidade"] = 99
	player := models.Player{
		Name:  "Sócrates",
		Stats: models.JSONB(stats),
	}
	err := repo.CreatePlayer(player)
	if err != nil {
		t.Fatalf("CreatePlayer() error = %v", err)
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
	db := setupTestDB(t)
	repo := NewPlayer(db)

	stats := make(map[string]int)
	stats["velocidade"] = 99
	player := models.Player{
		Name:  "Rivelino",
		Stats: models.JSONB(stats),
	}
	_ = repo.CreatePlayer(player)
	players := repo.GetPlayers()
	player = players[0]
	player.Name = "Rivelino Atualizado"

	err := repo.UpdatePlayer(player)
	if err != nil {
		t.Fatalf("UpdatePlayer() error = %v", err)
	}

	got, _ := repo.GetPlayerByID(player.ID)
	if got.Name != "Rivelino Atualizado" {
		t.Errorf("UpdatePlayer() = %v, want Name Rivelino Atualizado", got)
	}
}

func TestPlayerRepository_DeletePlayer(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPlayer(db)

	stats := make(map[string]int)
	stats["velocidade"] = 99
	player := models.Player{
		Name:  "Garrincha",
		Stats: models.JSONB(stats),
	}
	_ = repo.CreatePlayer(player)
	players := repo.GetPlayers()
	player = players[0]

	err := repo.DeletePlayer(player.ID)
	if err != nil {
		t.Fatalf("DeletePlayer() error = %v", err)
	}

	_, err = repo.GetPlayerByID(player.ID)
	if err == nil {
		t.Errorf("GetPlayerByID() after delete should return error, got nil")
	}
}
