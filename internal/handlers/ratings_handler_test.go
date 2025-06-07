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

type mockRatingService struct {
	services.RatingService
	createRatingFn  func(models.Rating) error
	getAllRatingsFn func() []models.Rating
	getRatingByIDFn func(int) (*models.Rating, error)
	updateRatingFn  func(models.Rating) error
	deleteRatingFn  func(int) error
}

func (m *mockRatingService) CreateRating(rating models.Rating) error {
	if m.createRatingFn != nil {
		return m.createRatingFn(rating)
	}
	return nil
}

func (m *mockRatingService) GetAllRatings() []models.Rating {
	if m.getAllRatingsFn != nil {
		return m.getAllRatingsFn()
	}
	return []models.Rating{}
}

func (m *mockRatingService) GetRatingByID(id int) (*models.Rating, error) {
	if m.getRatingByIDFn != nil {
		return m.getRatingByIDFn(id)
	}
	return nil, nil
}

func (m *mockRatingService) UpdateRating(rating models.Rating) error {
	if m.updateRatingFn != nil {
		return m.updateRatingFn(rating)
	}
	return nil
}

func (m *mockRatingService) DeleteRating(id int) error {
	if m.deleteRatingFn != nil {
		return m.deleteRatingFn(id)
	}
	return nil
}

func TestCreateRating(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        string
		mockService    services.RatingService
		expectedStatus int
	}{
		{
			name: "success",
			reqBody: `{
				"match_id": 1,
				"player_id": 1,
				"rated_player_id": 2,
				"finishing": 80,
				"passing": 75,
				"speed": 85,
				"defense": 70,
				"stamina": 90,
				"highlight": 95
			}`,
			mockService: &mockRatingService{
				createRatingFn: func(rating models.Rating) error {
					return nil
				},
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:    "invalid json",
			reqBody: `{"match_id": }`,
			mockService: &mockRatingService{
				createRatingFn: func(rating models.Rating) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &RatingHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodPost, "/ratings", strings.NewReader(tt.reqBody))
			rr := httptest.NewRecorder()

			handler.CreateRating(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestGetRatings(t *testing.T) {
	tests := []struct {
		name           string
		mockService    services.RatingService
		expectedStatus int
	}{
		{
			name: "success",
			mockService: &mockRatingService{
				getAllRatingsFn: func() []models.Rating {
					return []models.Rating{{MatchID: 1, PlayerID: 1}}
				},
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &RatingHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodGet, "/ratings", nil)
			rr := httptest.NewRecorder()

			handler.GetRatings(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.name == "success" {
				var ratings []models.Rating
				err := json.NewDecoder(rr.Body).Decode(&ratings)
				if err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}
				if len(ratings) != 1 {
					t.Errorf("Expected 1 rating, got %d", len(ratings))
				}
			}
		})
	}
}

func TestGetRatingByID(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockService    services.RatingService
		expectedStatus int
	}{
		{
			name: "success",
			id:   "1",
			mockService: &mockRatingService{
				getRatingByIDFn: func(id int) (*models.Rating, error) {
					return &models.Rating{MatchID: 1, PlayerID: 1}, nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "invalid id",
			id:   "invalid",
			mockService: &mockRatingService{
				getRatingByIDFn: func(id int) (*models.Rating, error) {
					return nil, nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &RatingHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodGet, "/ratings/"+tt.id, nil)
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.GetRatingByID(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestUpdateRating(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		reqBody        string
		mockService    services.RatingService
		expectedStatus int
	}{
		{
			name: "success",
			id:   "1",
			reqBody: `{
				"match_id": 1,
				"player_id": 1,
				"rated_player_id": 2,
				"finishing": 90,
				"passing": 85,
				"speed": 95,
				"defense": 80,
				"stamina": 95,
				"highlight": 100
			}`,
			mockService: &mockRatingService{
				updateRatingFn: func(rating models.Rating) error {
					return nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:    "invalid json",
			id:      "1",
			reqBody: `{"match_id": }`,
			mockService: &mockRatingService{
				updateRatingFn: func(rating models.Rating) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &RatingHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodPut, "/ratings/"+tt.id, strings.NewReader(tt.reqBody))
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.UpdateRating(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestDeleteRating(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockService    services.RatingService
		expectedStatus int
	}{
		{
			name: "success",
			id:   "1",
			mockService: &mockRatingService{
				deleteRatingFn: func(id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name: "invalid id",
			id:   "invalid",
			mockService: &mockRatingService{
				deleteRatingFn: func(id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &RatingHandler{Service: tt.mockService}
			req := httptest.NewRequest(http.MethodDelete, "/ratings/"+tt.id, nil)
			rr := httptest.NewRecorder()

			vars := map[string]string{"id": tt.id}
			req = mux.SetURLVars(req, vars)

			handler.DeleteRating(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}
