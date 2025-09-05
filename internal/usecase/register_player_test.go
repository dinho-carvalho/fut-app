package usecase

import (
	"errors"
	"testing"

	"fut-app/internal/domain"
	apperrors "fut-app/internal/errors"
)

// Mock implementation of RegisterPlayerGateway
type mockRegisterPlayerGateway struct {
	shouldReturnError bool
	returnedPlayer    *domain.Player
	returnedError     error
}

func (m *mockRegisterPlayerGateway) Register(player domain.Player) (*domain.Player, error) {
	if m.shouldReturnError {
		return nil, m.returnedError
	}
	return m.returnedPlayer, nil
}

func TestRegisterPlayerUseCase_Execute_Success(t *testing.T) {
	// Arrange
	expectedPlayer := &domain.Player{
		ID:   1,
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

	mockGateway := &mockRegisterPlayerGateway{
		shouldReturnError: false,
		returnedPlayer:    expectedPlayer,
	}

	useCase := NewPlayerUseCase(mockGateway)

	player := domain.Player{
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
	result, err := useCase.Execute(player)

	// Assert
	if err != nil {
		t.Errorf("Execute() error = %v, want nil", err)
	}

	if result == nil {
		t.Fatal("Execute() result = nil, want player")
	}

	if result.Name != expectedPlayer.Name {
		t.Errorf("Execute() name = %v, want %v", result.Name, expectedPlayer.Name)
	}

	if result.ID != expectedPlayer.ID {
		t.Errorf("Execute() ID = %v, want %v", result.ID, expectedPlayer.ID)
	}

	if len(result.Position) != len(expectedPlayer.Position) {
		t.Errorf("Execute() position count = %v, want %v", len(result.Position), len(expectedPlayer.Position))
	}
}

func TestRegisterPlayerUseCase_Execute_ValidationError(t *testing.T) {
	// Arrange
	mockGateway := &mockRegisterPlayerGateway{}
	useCase := NewPlayerUseCase(mockGateway)

	player := domain.Player{
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
	result, err := useCase.Execute(player)

	// Assert
	if err == nil {
		t.Fatal("Execute() error = nil, want validation error")
	}

	if result != nil {
		t.Errorf("Execute() result = %v, want nil", result)
	}

	// The validation error should be returned as-is from the domain
	_, ok := err.(*apperrors.ValidationErrors)
	if !ok {
		t.Fatalf("Execute() error type = %T, want *apperrors.ValidationErrors", err)
	}

	// Verify that the gateway was not called when validation fails
	if mockGateway.shouldReturnError {
		t.Error("Gateway should not have been called when validation fails")
	}
}

func TestRegisterPlayerUseCase_Execute_GatewayError(t *testing.T) {
	// Arrange
	expectedError := errors.New("database connection failed")
	mockGateway := &mockRegisterPlayerGateway{
		shouldReturnError: true,
		returnedError:     expectedError,
	}

	useCase := NewPlayerUseCase(mockGateway)

	player := domain.Player{
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
	result, err := useCase.Execute(player)

	// Assert
	if err == nil {
		t.Fatal("Execute() error = nil, want gateway error")
	}

	if result != nil {
		t.Errorf("Execute() result = %v, want nil", result)
	}

	if err != expectedError {
		t.Errorf("Execute() error = %v, want %v", err, expectedError)
	}
}

func TestRegisterPlayerUseCase_Execute_ValidPlayerWithMultiplePositions(t *testing.T) {
	// Arrange
	expectedPlayer := &domain.Player{
		ID:   1,
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

	mockGateway := &mockRegisterPlayerGateway{
		shouldReturnError: false,
		returnedPlayer:    expectedPlayer,
	}

	useCase := NewPlayerUseCase(mockGateway)

	player := domain.Player{
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
	result, err := useCase.Execute(player)

	// Assert
	if err != nil {
		t.Errorf("Execute() error = %v, want nil", err)
	}

	if result == nil {
		t.Fatal("Execute() result = nil, want player")
	}

	if len(result.Position) != 2 {
		t.Errorf("Execute() position count = %v, want 2", len(result.Position))
	}

	expectedPositions := []string{"Atacante", "Meio-campo"}
	for _, expectedPos := range expectedPositions {
		found := false
		for _, actualPos := range result.Position {
			if actualPos == expectedPos {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Execute() missing position %s, got %v", expectedPos, result.Position)
		}
	}
}
