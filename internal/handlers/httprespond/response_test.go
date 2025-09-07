package httprespond

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON_WritesStatusHeaderAndBody(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := map[string]string{"ok": "yes"}

	err := JSON(rr, http.StatusCreated, payload)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var got map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, "yes", got["ok"])
}

func TestJSON_DifferentStatusCodes(t *testing.T) {
	tests := []struct {
		name   string
		status int
	}{
		{"OK", http.StatusOK},
		{"Created", http.StatusCreated},
		{"BadRequest", http.StatusBadRequest},
		{"NotFound", http.StatusNotFound},
		{"InternalError", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			payload := map[string]string{"status": tt.name}

			err := JSON(rr, tt.status, payload)
			assert.NoError(t, err)
			assert.Equal(t, tt.status, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		})
	}
}

func TestJSON_DifferentDataTypes(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
	}{
		{"String", "test string"},
		{"Number", 42},
		{"Boolean", true},
		{"Array", []string{"a", "b", "c"}},
		{"Map", map[string]interface{}{"key": "value", "number": 123}},
		{"Nil", nil},
		{"Struct", struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{Name: "Test", Age: 25}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			err := JSON(rr, http.StatusOK, tt.data)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

			// Verify response is valid JSON
			var result interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &result)
			assert.NoError(t, err)
		})
	}
}

func TestJSON_EmptyData(t *testing.T) {
	rr := httptest.NewRecorder()

	err := JSON(rr, http.StatusNoContent, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, "null\n", rr.Body.String())
}

func TestJSON_ComplexStruct(t *testing.T) {
	type Player struct {
		ID        int            `json:"id"`
		Name      string         `json:"name"`
		Positions []string       `json:"positions"`
		Stats     map[string]int `json:"stats"`
	}

	player := Player{
		ID:        1,
		Name:      "Test Player",
		Positions: []string{"midfielder", "forward"},
		Stats:     map[string]int{"goals": 10, "assists": 5},
	}

	rr := httptest.NewRecorder()
	err := JSON(rr, http.StatusOK, player)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var result Player
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, player.ID, result.ID)
	assert.Equal(t, player.Name, result.Name)
	assert.Equal(t, player.Positions, result.Positions)
	assert.Equal(t, player.Stats, result.Stats)
}
