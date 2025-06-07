package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"fut-app/internal/database/models"
	"fut-app/internal/services"

	"github.com/gorilla/mux"
)

type mockMatchService struct {
	createMatchFn   func(models.Match) error
	getAllMatchesFn func() []models.Match
	getMatchByIDFn  func(int) (models.Match, error)
	updateMatchFn   func(models.Match) error
	deleteMatchFn   func(int) error
}

func (m *mockMatchService) CreateMatch(match models.Match) error {
	if m.createMatchFn != nil {
		return m.createMatchFn(match)
	}
	return nil
}

func (m *mockMatchService) GetAllMatches() []models.Match {
	if m.getAllMatchesFn != nil {
		return m.getAllMatchesFn()
	}
	return []models.Match{}
}

func (m *mockMatchService) GetMatchByID(id int) (models.Match, error) {
	if m.getMatchByIDFn != nil {
		return m.getMatchByIDFn(id)
	}
	return models.Match{}, nil
}

func (m *mockMatchService) UpdateMatch(match models.Match) error {
	if m.updateMatchFn != nil {
		return m.updateMatchFn(match)
	}
	return nil
}

func (m *mockMatchService) DeleteMatch(id int) error {
	if m.deleteMatchFn != nil {
		return m.deleteMatchFn(id)
	}
	return nil
}

func TestCreateMatch(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        string
		mockService    services.IMatchService
		expectedStatus int
	}{
		{
			name: "success",
			reqBody: `{
				"date": "2024-03-20T15:00:00Z",
				"location": "Campo 1",
				"team_a": ["1", "2", "3", "4", "5"],
				"team_b": ["6", "7", "8", "9", "10"],
				"score_a": 3,
				"score_b": 2
			}`,
			mockService: &mockMatchService{
				createMatchFn: func(match models.Match) error {
					return nil
				},
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:    "invalid json",
			reqBody: `{"date": }`,
			mockService: &mockMatchService{
				createMatchFn: func(match models.Match) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &MatchHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodPost, "/matches", strings.NewReader(tt.reqBody))
			rr := httptest.NewRecorder()

			handler.CreateMatch(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestGetMatches(t *testing.T) {
	tests := []struct {
		name           string
		mockService    services.IMatchService
		expectedStatus int
	}{
		{
			name: "success",
			mockService: &mockMatchService{
				getAllMatchesFn: func() []models.Match {
					return []models.Match{
						{
							Date:     time.Now(),
							Location: "Campo 1",
							TeamA:    []string{"1", "2", "3", "4", "5"},
							TeamB:    []string{"6", "7", "8", "9", "10"},
							ScoreA:   3,
							ScoreB:   2,
						},
					}
				},
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &MatchHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodGet, "/matches", nil)
			rr := httptest.NewRecorder()

			handler.GetMatches(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.name == "success" {
				var matches []models.Match
				err := json.NewDecoder(rr.Body).Decode(&matches)
				if err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}
				if len(matches) != 1 {
					t.Errorf("Expected 1 match, got %d", len(matches))
				}
			}
		})
	}
}

func TestGetMatchByID(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockService    services.IMatchService
		expectedStatus int
	}{
		{
			name: "success",
			id:   "1",
			mockService: &mockMatchService{
				getMatchByIDFn: func(id int) (models.Match, error) {
					return models.Match{
						Date:     time.Now(),
						Location: "Campo 1",
						TeamA:    []string{"1", "2", "3", "4", "5"},
						TeamB:    []string{"6", "7", "8", "9", "10"},
						ScoreA:   3,
						ScoreB:   2,
					}, nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "invalid id",
			id:   "invalid",
			mockService: &mockMatchService{
				getMatchByIDFn: func(id int) (models.Match, error) {
					return models.Match{}, nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &MatchHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodGet, "/matches/"+tt.id, nil)
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.GetMatchByID(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestUpdateMatch(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		reqBody        string
		mockService    services.IMatchService
		expectedStatus int
	}{
		{
			name: "success",
			id:   "1",
			reqBody: `{
				"date": "2024-03-20T15:00:00Z",
				"location": "Campo 2",
				"team_a": ["1", "2", "3", "4", "5"],
				"team_b": ["6", "7", "8", "9", "10"],
				"score_a": 4,
				"score_b": 2
			}`,
			mockService: &mockMatchService{
				updateMatchFn: func(match models.Match) error {
					return nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:    "invalid json",
			id:      "1",
			reqBody: `{"date": }`,
			mockService: &mockMatchService{
				updateMatchFn: func(match models.Match) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &MatchHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodPut, "/matches/"+tt.id, strings.NewReader(tt.reqBody))
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.UpdateMatch(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestDeleteMatch(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockService    services.IMatchService
		expectedStatus int
	}{
		{
			name: "success",
			id:   "1",
			mockService: &mockMatchService{
				deleteMatchFn: func(id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name: "invalid id",
			id:   "invalid",
			mockService: &mockMatchService{
				deleteMatchFn: func(id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &MatchHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodDelete, "/matches/"+tt.id, nil)
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.DeleteMatch(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}
