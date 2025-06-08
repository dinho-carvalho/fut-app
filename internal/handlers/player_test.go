package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"fut-app/internal/database/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type mockPlayerService struct {
	mock.Mock
}

func (m *mockPlayerService) CreatePlayer(p models.Player) error {
	args := m.Called(p)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (m *mockPlayerService) GetAllPlayers() []models.Player {
	args := m.Called()
	if players, ok := args.Get(0).([]models.Player); ok {
		return players
	}
	return nil
}

func (m *mockPlayerService) GetPlayerByID(id uint) (*models.Player, error) {
	args := m.Called(id)
	if player, ok := args.Get(0).(*models.Player); ok {
		return player, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockPlayerService) UpdatePlayer(p models.Player, id uint) error {
	args := m.Called(p, id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (m *mockPlayerService) DeletePlayer(id uint) error {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func TestPlayerHandler_GetPlayers(t *testing.T) {
	mockSvc := &mockPlayerService{}
	mockSvc.On("GetAllPlayers").Return([]models.Player{
		{Name: "Pelé"},
		{Name: "Garrincha"},
	})
	handler := &PlayerHandler{Service: mockSvc}
	req := httptest.NewRequest("GET", "/players", nil)
	w := httptest.NewRecorder()
	handler.GetPlayers(w, req)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Pelé")
	assert.Contains(t, w.Body.String(), "Garrincha")
}

func TestPlayerHandler_GetPlayerByID(t *testing.T) {
	mockSvc := &mockPlayerService{}
	mockSvc.On("GetPlayerByID", mock.Anything).Return(&models.Player{Name: "Pelé"}, nil)
	handler := &PlayerHandler{Service: mockSvc}

	// Sucesso
	req := httptest.NewRequest("GET", "/players/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	handler.GetPlayerByID(w, req)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Pelé")

	mockSvc.On("GetPlayerByID", mock.Anything).Return(nil, assert.AnError)

	// ID inválido
	req = httptest.NewRequest("GET", "/players/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})
	w = httptest.NewRecorder()
	handler.GetPlayerByID(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Invalid ID")
}

func TestPlayerHandler_UpdatePlayer(t *testing.T) {
	mockSvc := &mockPlayerService{}
	mockSvc.On("UpdatePlayer", mock.Anything, mock.Anything).Return(nil)
	handler := &PlayerHandler{Service: mockSvc}

	// Sucesso
	player := models.Player{Name: "Zico"}
	b, _ := json.Marshal(player)
	req := httptest.NewRequest("PUT", "/players/1", bytes.NewBuffer(b))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	handler.UpdatePlayer(w, req)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Zico")

	// ID inválido
	req = httptest.NewRequest("PUT", "/players/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})
	w = httptest.NewRecorder()
	handler.UpdatePlayer(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Invalid ID")

	// JSON inválido
	req = httptest.NewRequest("PUT", "/players/1", bytes.NewBuffer([]byte("{invalid")))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w = httptest.NewRecorder()
	handler.UpdatePlayer(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Invalid JSON")
}

func TestPlayerHandler_DeletePlayer(t *testing.T) {
	mockSvc := &mockPlayerService{}
	mockSvc.On("DeletePlayer", mock.Anything).Return(nil)
	handler := &PlayerHandler{Service: mockSvc}

	// Sucesso
	req := httptest.NewRequest("DELETE", "/players/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	handler.DeletePlayer(w, req)
	assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Player deleted successfully")

	// ID inválido
	req = httptest.NewRequest("DELETE", "/players/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})
	w = httptest.NewRecorder()
	handler.DeletePlayer(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Invalid ID")
}
