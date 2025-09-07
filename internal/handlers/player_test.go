package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fut-app/internal/domain"
	appErrors "fut-app/internal/errors"
	"fut-app/internal/handlers/dto"
)

// stubRegisterPlayerUseCase is a simple stub implementing RegisterPlayerUseCase
type stubRegisterPlayerUseCase struct {
	executeFn func(domain.Player) (*domain.Player, error)
}

func (s *stubRegisterPlayerUseCase) Execute(p domain.Player) (*domain.Player, error) {
	return s.executeFn(p)
}

func TestPlayerHandler_CreatePlayer_Success(t *testing.T) {
	// Arrange
	input := dto.PlayerDTO{
		Name: "Neymar Jr",
		Stats: map[string]interface{}{
			"pace": 90, "shooting": 85, "passing": 86, "dribbling": 93, "defending": 32, "physical": 58,
		},
		Position: []string{"LW"},
	}

	expected := &domain.Player{
		ID:       1,
		Name:     input.Name,
		Stats:    input.Stats,
		Position: input.Position,
	}

	uc := &stubRegisterPlayerUseCase{
		executeFn: func(p domain.Player) (*domain.Player, error) {
			// Ensure DTO was converted correctly
			if p.Name != input.Name {
				t.Fatalf("expected name %q, got %q", input.Name, p.Name)
			}
			return expected, nil
		},
	}

	h := NewPlayerHandler(uc)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/players", nil)

	// Act
	err := h.CreatePlayer(rr, req, input)
	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected content-type application/json, got %q", ct)
	}

	var got domain.Player
	if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}
	if got.ID != expected.ID || got.Name != expected.Name {
		t.Fatalf("unexpected body: %+v", got)
	}
	if len(got.Position) != len(expected.Position) || got.Position[0] != expected.Position[0] {
		t.Fatalf("unexpected positions: %+v", got.Position)
	}
	if len(got.Stats) != len(expected.Stats) {
		t.Fatalf("unexpected stats length: %d", len(got.Stats))
	}
}

func TestPlayerHandler_CreatePlayer_ErrorFromUseCase(t *testing.T) {
	// Arrange
	input := dto.PlayerDTO{
		Name:     "",
		Stats:    map[string]interface{}{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6},
		Position: []string{"LW"},
	}

	uc := &stubRegisterPlayerUseCase{
		executeFn: func(p domain.Player) (*domain.Player, error) {
			return nil, appErrors.ErrInvalidData
		},
	}

	h := NewPlayerHandler(uc)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/players", nil)

	// Act
	err := h.CreatePlayer(rr, req, input)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != appErrors.ErrInvalidData {
		t.Fatalf("expected ErrInvalidData, got %v", err)
	}
	// Handler should not have written a response on error
	if rr.Body.Len() != 0 {
		t.Fatalf("expected empty body on error, got %q", rr.Body.String())
	}
}
