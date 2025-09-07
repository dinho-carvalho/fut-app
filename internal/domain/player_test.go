package domain

import (
	"testing"

	"fut-app/internal/errors"
)

func TestPlayer_Validate_Success(t *testing.T) {
	// Arrange
	player := Player{
		Name: "Pelé",
		Stats: map[string]interface{}{
			"velocidade":  99,
			"drible":      95,
			"finalizacao": 90,
			"passe":       88,
			"defesa":      60,
			"fisico":      85,
		},
		Position: []string{"Atacante"},
	}

	// Act
	err := player.Validate()
	// Assert
	if err != nil {
		t.Errorf("Validate() error = %v, want nil", err)
	}
}

func TestPlayer_Validate_EmptyName(t *testing.T) {
	// Arrange
	player := Player{
		Name: "", // Empty name should cause validation error
		Stats: map[string]interface{}{
			"velocidade":  99,
			"drible":      95,
			"finalizacao": 90,
			"passe":       88,
			"defesa":      60,
			"fisico":      85,
		},
		Position: []string{"Atacante"},
	}

	// Act
	err := player.Validate()

	// Assert
	if err == nil {
		t.Fatal("Validate() error = nil, want validation error")
	}

	validationErr, ok := err.(*errors.ValidationErrors)
	if !ok {
		t.Fatalf("Validate() error type = %T, want *errors.ValidationErrors", err)
	}

	if len(*validationErr) != 1 {
		t.Errorf("Validate() validation errors count = %v, want 1", len(*validationErr))
	}

	if (*validationErr)[0].Field != "name" {
		t.Errorf("Validate() validation error field = %v, want 'name'", (*validationErr)[0].Field)
	}

	if (*validationErr)[0].Message != "Name is required" {
		t.Errorf("Validate() validation error message = %v, want 'Name is required'", (*validationErr)[0].Message)
	}
}

func TestPlayer_Validate_WrongStatsCount(t *testing.T) {
	// Arrange
	player := Player{
		Name: "Pelé",
		Stats: map[string]interface{}{
			"velocidade": 99,
			"drible":     95,
			// Only 2 stats instead of 6
		},
		Position: []string{"Atacante"},
	}

	// Act
	err := player.Validate()

	// Assert
	if err == nil {
		t.Fatal("Validate() error = nil, want validation error")
	}

	validationErr, ok := err.(*errors.ValidationErrors)
	if !ok {
		t.Fatalf("Validate() error type = %T, want *errors.ValidationErrors", err)
	}

	if len(*validationErr) != 1 {
		t.Errorf("Validate() validation errors count = %v, want 1", len(*validationErr))
	}

	if (*validationErr)[0].Field != "stats" {
		t.Errorf("Validate() validation error field = %v, want 'stats'", (*validationErr)[0].Field)
	}

	if (*validationErr)[0].Message != "Stats must contain exactly 6 keys" {
		t.Errorf("Validate() validation error message = %v, want 'Stats must contain exactly 6 keys'", (*validationErr)[0].Message)
	}
}

func TestPlayer_Validate_EmptyPositions(t *testing.T) {
	// Arrange
	player := Player{
		Name: "Pelé",
		Stats: map[string]interface{}{
			"velocidade":  99,
			"drible":      95,
			"finalizacao": 90,
			"passe":       88,
			"defesa":      60,
			"fisico":      85,
		},
		Position: []string{}, // Empty positions should cause validation error
	}

	// Act
	err := player.Validate()

	// Assert
	if err == nil {
		t.Fatal("Validate() error = nil, want validation error")
	}

	validationErr, ok := err.(*errors.ValidationErrors)
	if !ok {
		t.Fatalf("Validate() error type = %T, want *errors.ValidationErrors", err)
	}

	if len(*validationErr) != 1 {
		t.Errorf("Validate() validation errors count = %v, want 1", len(*validationErr))
	}

	if (*validationErr)[0].Field != "positions" {
		t.Errorf("Validate() validation error field = %v, want 'positions'", (*validationErr)[0].Field)
	}

	if (*validationErr)[0].Message != "At least one position is required" {
		t.Errorf("Validate() validation error message = %v, want 'At least one position is required'", (*validationErr)[0].Message)
	}
}

func TestPlayer_Validate_MultipleErrors(t *testing.T) {
	// Arrange
	player := Player{
		Name: "", // Empty name
		Stats: map[string]interface{}{
			"velocidade": 99,
			// Only 1 stat instead of 6
		},
		Position: []string{}, // Empty positions
	}

	// Act
	err := player.Validate()

	// Assert
	if err == nil {
		t.Fatal("Validate() error = nil, want validation error")
	}

	validationErr, ok := err.(*errors.ValidationErrors)
	if !ok {
		t.Fatalf("Validate() error type = %T, want *errors.ValidationErrors", err)
	}

	if len(*validationErr) != 3 {
		t.Errorf("Validate() validation errors count = %v, want 3", len(*validationErr))
	}

	// Check that all three validation errors are present
	fields := make(map[string]bool)
	for _, ve := range *validationErr {
		fields[ve.Field] = true
	}

	expectedFields := []string{"name", "stats", "positions"}
	for _, field := range expectedFields {
		if !fields[field] {
			t.Errorf("Validate() missing validation error for field: %s", field)
		}
	}
}

func TestPlayer_Validate_ValidPlayerWithMultiplePositions(t *testing.T) {
	// Arrange
	player := Player{
		Name: "Pelé",
		Stats: map[string]interface{}{
			"velocidade":  99,
			"drible":      95,
			"finalizacao": 90,
			"passe":       88,
			"defesa":      60,
			"fisico":      85,
		},
		Position: []string{"Atacante", "Meio-campo"},
	}

	// Act
	err := player.Validate()
	// Assert
	if err != nil {
		t.Errorf("Validate() error = %v, want nil", err)
	}
}

func TestPlayer_Validate_ExactStatsCount(t *testing.T) {
	// Arrange
	player := Player{
		Name: "Pelé",
		Stats: map[string]interface{}{
			"velocidade":  99,
			"drible":      95,
			"finalizacao": 90,
			"passe":       88,
			"defesa":      60,
			"fisico":      85,
		},
		Position: []string{"Atacante"},
	}

	// Act
	err := player.Validate()
	// Assert
	if err != nil {
		t.Errorf("Validate() error = %v, want nil", err)
	}
}

func TestPlayer_Validate_TooManyStats(t *testing.T) {
	// Arrange
	player := Player{
		Name: "Pelé",
		Stats: map[string]interface{}{
			"velocidade":  99,
			"drible":      95,
			"finalizacao": 90,
			"passe":       88,
			"defesa":      60,
			"fisico":      85,
			"extra":       100, // 7 stats instead of 6
		},
		Position: []string{"Atacante"},
	}

	// Act
	err := player.Validate()

	// Assert
	if err == nil {
		t.Fatal("Validate() error = nil, want validation error")
	}

	validationErr, ok := err.(*errors.ValidationErrors)
	if !ok {
		t.Fatalf("Validate() error type = %T, want *errors.ValidationErrors", err)
	}

	if len(*validationErr) != 1 {
		t.Errorf("Validate() validation errors count = %v, want 1", len(*validationErr))
	}

	if (*validationErr)[0].Field != "stats" {
		t.Errorf("Validate() validation error field = %v, want 'stats'", (*validationErr)[0].Field)
	}
}

func TestNewPlayer(t *testing.T) {
	// Arrange
	name := "Pelé"
	stats := map[string]interface{}{
		"velocidade":  99,
		"drible":      95,
		"finalizacao": 90,
		"passe":       88,
		"defesa":      60,
		"fisico":      85,
	}
	positions := []string{"Atacante", "Meio-campo"}

	// Act
	player := NewPlayer(name, stats, positions)

	// Assert
	if player == nil {
		t.Fatal("NewPlayer() returned nil")
	}

	if player.Name != name {
		t.Errorf("NewPlayer() name = %v, want %v", player.Name, name)
	}

	if len(player.Stats) != len(stats) {
		t.Errorf("NewPlayer() stats length = %v, want %v", len(player.Stats), len(stats))
	}

	if len(player.Position) != len(positions) {
		t.Errorf("NewPlayer() position length = %v, want %v", len(player.Position), len(positions))
	}

	// Check if all stats are present
	for key, value := range stats {
		if player.Stats[key] != value {
			t.Errorf("NewPlayer() stats[%s] = %v, want %v", key, player.Stats[key], value)
		}
	}

	// Check if all positions are present
	for i, pos := range positions {
		if player.Position[i] != pos {
			t.Errorf("NewPlayer() position[%d] = %v, want %v", i, player.Position[i], pos)
		}
	}
}
