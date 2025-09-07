package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	appErrors "fut-app/internal/errors"
)

func TestAppHandler_ServesHTTP_ValidationErrors(t *testing.T) {
	h := AppHandler(func(w http.ResponseWriter, r *http.Request) error {
		var ve appErrors.ValidationErrors
		ve.Append("field", "msg")
		return &ve
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestAppHandler_ServesHTTP_MapsKnownErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{"not_found", appErrors.ErrNotFound, http.StatusNotFound},
		{"bad_request", appErrors.ErrBadRequest, http.StatusBadRequest},
		{"already_exists", appErrors.ErrAlreadyExists, http.StatusConflict},
		{"invalid_data", appErrors.ErrInvalidData, http.StatusUnprocessableEntity},
		{"unauthorized", appErrors.ErrUnauthorized, http.StatusUnauthorized},
		{"forbidden", appErrors.ErrForbidden, http.StatusForbidden},
		{"database", appErrors.ErrDatabase, http.StatusInternalServerError},
		{"internal", errors.New("other"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AppHandler(func(w http.ResponseWriter, r *http.Request) error { return tt.err })
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			h.ServeHTTP(rr, req)
			if rr.Code != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, rr.Code)
			}
		})
	}
}
