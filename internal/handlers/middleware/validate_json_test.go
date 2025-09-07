package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

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
		assert.Equal(t, dto.Name, d.Name)
		assert.Equal(t, len(dto.Stats), len(d.Stats))
		// JSON unmarshaling converts numbers to float64, so we verify structure instead of exact equality
		for key, expectedVal := range dto.Stats {
			actualVal, exists := d.Stats[key]
			assert.True(t, exists, "Key %s should exist", key)
			assert.Equal(t, float64(expectedVal.(int)), actualVal, "Value for key %s should match", key)
		}
		return nil
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	err := handler(rr, req)
	assert.NoError(t, err)
	assert.True(t, called)
}

func TestValidateJSON_InvalidBody(t *testing.T) {
	body := []byte(`{"name": "p", "stats": 123}`) // wrong type for stats
	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error { return nil })

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	err := handler(rr, req)
	assert.Error(t, err)
	assert.Equal(t, appErrors.ErrInvalidData, err)
}

func TestValidateJSON_ValidationFails(t *testing.T) {
	dto := sampleDTO{Name: "", Stats: map[string]interface{}{"a": 1, "b": 2}} // invalid: required + statslen
	b, _ := json.Marshal(dto)

	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error { return nil })
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	err := handler(rr, req)
	assert.Error(t, err)
	assert.Equal(t, appErrors.ErrInvalidData, err)
}

func TestValidateJSON_InvalidJSON(t *testing.T) {
	body := []byte(`{"name": "incomplete json"`) // malformed JSON
	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error { return nil })

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	err := handler(rr, req)
	assert.Error(t, err)
	assert.Equal(t, appErrors.ErrInvalidData, err)
}

func TestValidateJSON_EmptyBody(t *testing.T) {
	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error { return nil })

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte{}))
	err := handler(rr, req)
	assert.Error(t, err)
	assert.Equal(t, appErrors.ErrInvalidData, err)
}

func TestValidateJSON_UnknownFields(t *testing.T) {
	body := []byte(`{"name": "Player", "stats": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}, "unknown": "field"}`)
	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error { return nil })

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	err := handler(rr, req)
	assert.Error(t, err)
	assert.Equal(t, appErrors.ErrInvalidData, err)
}

func TestValidateJSON_CustomValidation_statslen(t *testing.T) {
	tests := []struct {
		name      string
		statsLen  int
		shouldErr bool
	}{
		{"exactly 6 stats - valid", 6, false},
		{"less than 6 stats - invalid", 3, true},
		{"more than 6 stats - invalid", 8, true},
		{"empty stats - invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats := make(map[string]interface{})
			for i := 0; i < tt.statsLen; i++ {
				stats[string(rune('a'+i))] = i + 1
			}

			dto := sampleDTO{
				Name:  "Valid Name",
				Stats: stats,
			}
			b, _ := json.Marshal(dto)

			handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error { return nil })
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
			err := handler(rr, req)

			if tt.shouldErr {
				assert.Error(t, err)
				assert.Equal(t, appErrors.ErrInvalidData, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

type anotherDTO struct {
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,min=0,max=120"`
}

func TestValidateJSON_DifferentStructTypes(t *testing.T) {
	dto := anotherDTO{
		Email: "test@example.com",
		Age:   25,
	}
	b, _ := json.Marshal(dto)

	called := false
	handler := ValidateJSON[anotherDTO](func(w http.ResponseWriter, r *http.Request, d anotherDTO) error {
		called = true
		assert.Equal(t, dto.Email, d.Email)
		assert.Equal(t, dto.Age, d.Age)
		return nil
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	err := handler(rr, req)
	assert.NoError(t, err)
	assert.True(t, called)
}

func TestValidateJSON_NextFunctionError(t *testing.T) {
	dto := sampleDTO{
		Name:  "Valid Name",
		Stats: map[string]interface{}{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6},
	}
	b, _ := json.Marshal(dto)

	expectedError := appErrors.ErrDatabase
	handler := ValidateJSON[sampleDTO](func(w http.ResponseWriter, r *http.Request, d sampleDTO) error {
		return expectedError
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	err := handler(rr, req)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
