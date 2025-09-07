package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	appErrors "fut-app/internal/errors"
)

type sampleDTO struct {
	Name  string                 `json:"name" validate:"required"`
	Stats map[string]interface{} `json:"stats" validate:"required,statslen"`
}

func TestValidateJSON_Success(t *testing.T) {
	dto := sampleDTO{
		Name:  "Player",
		Stats: map[string]interface{}{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6},
	}
	b, _ := json.Marshal(dto)

	called := false
	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error {
		called = true
		if d.Name != dto.Name {
			t.Fatalf("unexpected dto: %+v", d)
		}
		return nil
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	err := handler(rr, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !called {
		t.Fatalf("expected next to be called")
	}
}

func TestValidateJSON_InvalidBody(t *testing.T) {
	body := []byte(`{"name": "p", "stats": 123}`) // wrong type for stats
	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error { return nil })

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	err := handler(rr, req)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != appErrors.ErrInvalidData {
		t.Fatalf("expected ErrInvalidData, got %v", err)
	}
}

func TestValidateJSON_ValidationFails(t *testing.T) {
	dto := sampleDTO{Name: "", Stats: map[string]interface{}{"a": 1, "b": 2}} // invalid: required + statslen
	b, _ := json.Marshal(dto)

	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error { return nil })
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	err := handler(rr, req)
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}
	if err != appErrors.ErrInvalidData {
		t.Fatalf("expected ErrInvalidData, got %v", err)
	}
}
