package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"fut-app/internal/database/models"
	"fut-app/internal/services"

	"github.com/gorilla/mux"
)

type mockPlayerService struct {
	services.PlayerService
	createPlayerFn  func(models.Player) error
	getAllPlayersFn func() []models.Player
	getPlayerByIDFn func(int) (*models.Player, error)
	updatePlayerFn  func(models.Player, int) error
	deletePlayerFn  func(int) error
}

func (m *mockPlayerService) CreatePlayer(player models.Player) error {
	if m.createPlayerFn != nil {
		return m.createPlayerFn(player)
	}
	return nil
}

func (m *mockPlayerService) GetAllPlayers() []models.Player {
	if m.getAllPlayersFn != nil {
		return m.getAllPlayersFn()
	}
	return []models.Player{}
}

func (m *mockPlayerService) GetPlayerByID(id int) (*models.Player, error) {
	if m.getPlayerByIDFn != nil {
		return m.getPlayerByIDFn(id)
	}
	return nil, nil
}

func (m *mockPlayerService) UpdatePlayer(player models.Player, id int) error {
	if m.updatePlayerFn != nil {
		return m.updatePlayerFn(player, id)
	}
	return nil
}

func (m *mockPlayerService) DeletePlayer(id int) error {
	if m.deletePlayerFn != nil {
		return m.deletePlayerFn(id)
	}
	return nil
}

func TestCreatePlayer(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        string
		mockService    services.PlayerService
		expectedStatus int
	}{
		{
			name:    "success",
			reqBody: `{"name": "Test Player"}`,
			mockService: &mockPlayerService{
				createPlayerFn: func(player models.Player) error {
					return nil
				},
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:    "invalid json",
			reqBody: `{"name": }`,
			mockService: &mockPlayerService{
				createPlayerFn: func(player models.Player) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &PlayerHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodPost, "/players", strings.NewReader(tt.reqBody))
			rr := httptest.NewRecorder()

			handler.CreatePlayer(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestGetPlayers(t *testing.T) {
	tests := []struct {
		name           string
		mockService    services.PlayerService
		expectedStatus int
	}{
		{
			name: "success",
			mockService: &mockPlayerService{
				getAllPlayersFn: func() []models.Player {
					return []models.Player{{Name: "Test Player"}}
				},
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &PlayerHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodGet, "/players", nil)
			rr := httptest.NewRecorder()

			handler.GetPlayers(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestGetPlayerByID(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockService    services.PlayerService
		expectedStatus int
	}{
		{
			name: "success",
			id:   "1",
			mockService: &mockPlayerService{
				getPlayerByIDFn: func(id int) (*models.Player, error) {
					return &models.Player{Name: "Test Player"}, nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "invalid id",
			id:   "invalid",
			mockService: &mockPlayerService{
				getPlayerByIDFn: func(id int) (*models.Player, error) {
					return nil, nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &PlayerHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodGet, "/players/"+tt.id, nil)
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.GetPlayerByID(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestUpdatePlayer(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		reqBody        string
		mockService    services.PlayerService
		expectedStatus int
	}{
		{
			name:    "success",
			id:      "1",
			reqBody: `{"name": "Updated Player"}`,
			mockService: &mockPlayerService{
				updatePlayerFn: func(player models.Player, id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:    "invalid json",
			id:      "1",
			reqBody: `{"name": }`,
			mockService: &mockPlayerService{
				updatePlayerFn: func(player models.Player, id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &PlayerHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodPut, "/players/"+tt.id, strings.NewReader(tt.reqBody))
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.UpdatePlayer(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestDeletePlayer(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockService    services.PlayerService
		expectedStatus int
	}{
		{
			name: "success",
			id:   "1",
			mockService: &mockPlayerService{
				deletePlayerFn: func(id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name: "invalid id",
			id:   "invalid",
			mockService: &mockPlayerService{
				deletePlayerFn: func(id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &PlayerHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodDelete, "/players/"+tt.id, nil)
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.DeletePlayer(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}
