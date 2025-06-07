package services

import (
	"testing"

	"fut-app/internal/database/models"
	// "fut-app/internal/repositories" // Not strictly needed for type compilation with nil repo
)

// TestMatchService_CreateMatch tests CreateMatch with a nil repository.
func TestMatchService_CreateMatch(t *testing.T) {
	service := NewMatchService(nil) // Initialize with nil repository
	match := models.Match{}        // Zero-value model

	defer func() {
		if r := recover(); r != nil {
			t.Logf("CreateMatch panicked as expected with nil repo: %v", r)
		} else {
			// This path should ideally not be hit if the method directly calls the repo.
			// If an error is returned before panic, this indicates different behavior.
			t.Errorf("CreateMatch did not panic with nil repo, check for preliminary error return or logic.")
		}
	}()

	// Call the method, expecting a panic
	err := service.CreateMatch(match)

	// This part of the test might only be reached if the method returns an error *before*
	// trying to use the nil repository, which is not the case for current service structure.
	// If an error is returned (e.g. from validation) and no panic, the defer func won't log a panic.
	if err == nil && recover() == nil { // Check if no error returned AND no panic occurred
		t.Errorf("CreateMatch: expected panic or error with nil repo, got nil error and no panic")
	}
	t.Log("TestMatchService_CreateMatch executed (placeholder, expecting panic)")
}

// TestMatchService_GetAllMatches tests GetAllMatches with a nil repository.
func TestMatchService_GetAllMatches(t *testing.T) {
	service := NewMatchService(nil)

	defer func() {
		if r := recover(); r != nil {
			t.Logf("GetAllMatches panicked as expected with nil repo: %v", r)
		} else {
			t.Errorf("GetAllMatches did not panic with nil repo")
		}
	}()

	_ = service.GetAllMatches() // Expect panic
	t.Log("TestMatchService_GetAllMatches executed (placeholder, expecting panic)")
}

// TestMatchService_GetMatchByID tests GetMatchByID with a nil repository.
func TestMatchService_GetMatchByID(t *testing.T) {
	service := NewMatchService(nil)
	id := 1

	defer func() {
		if r := recover(); r != nil {
			t.Logf("GetMatchByID panicked as expected with nil repo: %v", r)
		} else {
			t.Errorf("GetMatchByID did not panic with nil repo")
		}
	}()

	_, err := service.GetMatchByID(id) // Expect panic
	if err == nil && recover() == nil {
		t.Errorf("GetMatchByID: expected panic or error with nil repo, got nil error and no panic")
	}
	t.Log("TestMatchService_GetMatchByID executed (placeholder, expecting panic)")
}

// TestMatchService_UpdateMatch tests UpdateMatch with a nil repository.
func TestMatchService_UpdateMatch(t *testing.T) {
	service := NewMatchService(nil)
	match := models.Match{ID: 1}

	defer func() {
		if r := recover(); r != nil {
			t.Logf("UpdateMatch panicked as expected with nil repo: %v", r)
		} else {
			t.Errorf("UpdateMatch did not panic with nil repo")
		}
	}()

	err := service.UpdateMatch(match) // Expect panic
	if err == nil && recover() == nil {
		t.Errorf("UpdateMatch: expected panic or error with nil repo, got nil error and no panic")
	}
	t.Log("TestMatchService_UpdateMatch executed (placeholder, expecting panic)")
}

// TestMatchService_DeleteMatch tests DeleteMatch with a nil repository.
func TestMatchService_DeleteMatch(t *testing.T) {
	service := NewMatchService(nil)
	id := 1

	defer func() {
		if r := recover(); r != nil {
			t.Logf("DeleteMatch panicked as expected with nil repo: %v", r)
		} else {
			t.Errorf("DeleteMatch did not panic with nil repo")
		}
	}()

	err := service.DeleteMatch(id) // Expect panic
	if err == nil && recover() == nil {
		t.Errorf("DeleteMatch: expected panic or error with nil repo, got nil error and no panic")
	}
	t.Log("TestMatchService_DeleteMatch executed (placeholder, expecting panic)")
}
