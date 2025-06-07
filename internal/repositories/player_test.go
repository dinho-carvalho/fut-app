package repositories

import (
	"testing"

	"fut-app/internal/database"

	"fut-app/internal/database/models"
)

// TestPlayerRepository_CreatePlayer tests the CreatePlayer method with a nil DB.
func TestPlayerRepository_CreatePlayer(t *testing.T) {
	tests := []struct {
		name    string
		player  models.Player
		wantErr bool
	}{
		{
			name: "success",
			player: models.Player{
				Name: "John Doe",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			player: models.Player{
				Name: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupTestDB(t)
			repo := NewPlayer(db)

			err := repo.CreatePlayer(tt.player)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var found models.Player
				result := db.First(&found)
				if result.Error != nil {
					t.Errorf("Failed to find created player: %v", result.Error)
				}
				if found.Name != tt.player.Name {
					t.Errorf("Created player name = %v, want %v", found.Name, tt.player.Name)
				}
			}
		})
	}
}

// TestPlayerRepository_GetPlayers tests the GetPlayers method with a nil DB.
func TestPlayerRepository_GetPlayers(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPlayer(db)

	// Create test players
	testPlayers := []models.Player{
		{Name: "Player 1"},
		{Name: "Player 2"},
		{Name: "Player 3"},
	}
	for _, p := range testPlayers {
		if err := db.Create(&p).Error; err != nil {
			t.Fatalf("Failed to create test player: %v", err)
		}
	}

	got := repo.GetPlayers()
	if len(got) != len(testPlayers) {
		t.Errorf("GetPlayers() got = %v players, want %v", len(got), len(testPlayers))
	}
}

// TestPlayerRepository_GetPlayerByID tests the GetPlayerByID method with a nil DB.
func TestPlayerRepository_GetPlayerByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPlayer(db)

	// Create a test player
	testPlayer := models.Player{Name: "Test Player"}
	if err := db.Create(&testPlayer).Error; err != nil {
		t.Fatalf("Failed to create test player: %v", err)
	}

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "existing player",
			id:      int(testPlayer.ID),
			wantErr: false,
		},
		{
			name:    "non-existing player",
			id:      9999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player, err := repo.GetPlayerByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && player.Name != testPlayer.Name {
				t.Errorf("GetPlayerByID() got = %v, want %v", player.Name, testPlayer.Name)
			}
		})
	}
}

// TestPlayerRepository_UpdatePlayer tests the UpdatePlayer method with a nil DB.
func TestPlayerRepository_UpdatePlayer(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPlayer(db)

	// Create a test player
	testPlayer := models.Player{Name: "Test Player"}
	if err := db.Create(&testPlayer).Error; err != nil {
		t.Fatalf("Failed to create test player: %v", err)
	}

	tests := []struct {
		name    string
		player  models.Player
		wantErr bool
	}{
		{
			name: "valid update",
			player: models.Player{
				Model: database.Model{
					ID: testPlayer.ID,
				},
				Name: "Updated Player",
			},
			wantErr: false,
		},
		{
			name: "non-existing player",
			player: models.Player{
				Model: database.Model{
					ID: 9999,
				},
				Name: "Non-existent Player",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdatePlayer(tt.player)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var updated models.Player
				if err := db.First(&updated, tt.player.ID).Error; err != nil {
					t.Errorf("Failed to find updated player: %v", err)
				}
				if updated.Name != tt.player.Name {
					t.Errorf("Updated player name = %v, want %v", updated.Name, tt.player.Name)
				}
			}
		})
	}
}

// TestPlayerRepository_DeletePlayer tests the DeletePlayer method with a nil DB.
func TestPlayerRepository_DeletePlayer(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPlayer(db)

	// Create a test player
	testPlayer := models.Player{Name: "Test Player"}
	if err := db.Create(&testPlayer).Error; err != nil {
		t.Fatalf("Failed to create test player: %v", err)
	}

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "existing player",
			id:      int(testPlayer.ID),
			wantErr: false,
		},
		{
			name:    "non-existing player",
			id:      9999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.DeletePlayer(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var found models.Player
				err := db.First(&found, tt.id).Error
				if err == nil {
					t.Error("DeletePlayer() player still exists after deletion")
				}
			}
		})
	}
}
