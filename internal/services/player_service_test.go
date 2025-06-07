package services

import (
	"testing"

	"fut-app/internal/database/models"
	// "fut-app/internal/repositories" // Not strictly needed for type compilation with nil repo
)

// TestPlayerService_CreatePlayer tests CreatePlayer with a nil repository.
func TestPlayerService_CreatePlayer(t *testing.T) {
	service := NewPlayerService(nil) // Initialize with nil repository
	player := models.Player{}       // Zero-value model

	defer func() {
		if r := recover(); r != nil {
			t.Logf("CreatePlayer panicked as expected with nil repo: %v", r)
		} else {
			t.Errorf("CreatePlayer did not panic with nil repo")
		}
	}()

	err := service.CreatePlayer(player) // Expect panic
	if err == nil && recover() == nil {
		t.Errorf("CreatePlayer: expected panic or error with nil repo, got nil error and no panic")
	}
	t.Log("TestPlayerService_CreatePlayer executed (placeholder, expecting panic)")
}

// TestPlayerService_GetAllPlayers tests GetAllPlayers with a nil repository.
func TestPlayerService_GetAllPlayers(t *testing.T) {
	service := NewPlayerService(nil)

	defer func() {
		if r := recover(); r != nil {
			t.Logf("GetAllPlayers panicked as expected with nil repo: %v", r)
		} else {
			t.Errorf("GetAllPlayers did not panic with nil repo")
		}
	}()

	_ = service.GetAllPlayers() // Expect panic
	t.Log("TestPlayerService_GetAllPlayers executed (placeholder, expecting panic)")
}

// TestPlayerService_GetPlayerByID tests GetPlayerByID with a nil repository.
func TestPlayerService_GetPlayerByID(t *testing.T) {
	service := NewPlayerService(nil)
	id := 1

	defer func() {
		if r := recover(); r != nil {
			t.Logf("GetPlayerByID panicked as expected with nil repo: %v", r)
		} else {
			t.Errorf("GetPlayerByID did not panic with nil repo")
		}
	}()

	_, err := service.GetPlayerByID(id) // Expect panic
	if err == nil && recover() == nil {
		t.Errorf("GetPlayerByID: expected panic or error with nil repo, got nil error and no panic")
	}
	t.Log("TestPlayerService_GetPlayerByID executed (placeholder, expecting panic)")
}

// TestPlayerService_UpdatePlayer tests UpdatePlayer with a nil repository.
// This should panic when calling GetPlayerByID on the nil repo.
func TestPlayerService_UpdatePlayer(t *testing.T) {
	service := NewPlayerService(nil)
	player := models.Player{ID: 1}
	id := 1

	defer func() {
		if r := recover(); r != nil {
			t.Logf("UpdatePlayer panicked as expected with nil repo (on GetPlayerByID call): %v", r)
		} else {
			t.Errorf("UpdatePlayer did not panic with nil repo")
		}
	}()

	err := service.UpdatePlayer(player, id) // Expect panic
	if err == nil && recover() == nil {
		t.Errorf("UpdatePlayer: expected panic or error with nil repo, got nil error and no panic")
	}
	t.Log("TestPlayerService_UpdatePlayer executed (placeholder, expecting panic)")
}

// TestPlayerService_DeletePlayer tests DeletePlayer with a nil repository.
func TestPlayerService_DeletePlayer(t *testing.T) {
	service := NewPlayerService(nil)
	id := 1

	defer func() {
		if r := recover(); r != nil {
			t.Logf("DeletePlayer panicked as expected with nil repo: %v", r)
		} else {
			t.Errorf("DeletePlayer did not panic with nil repo")
		}
	}()

	err := service.DeletePlayer(id) // Expect panic
	if err == nil && recover() == nil {
		t.Errorf("DeletePlayer: expected panic or error with nil repo, got nil error and no panic")
	}
	t.Log("TestPlayerService_DeletePlayer executed (placeholder, expecting panic)")
}
