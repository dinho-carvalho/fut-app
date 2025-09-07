package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

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

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "bad_request", response["code"])
	assert.Equal(t, "Validation failed", response["message"])
	assert.NotNil(t, response["errors"])
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
			assert.Equal(t, tt.want, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

			var response appErrors.HTTPError
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.NotEmpty(t, response.Code)
			assert.NotEmpty(t, response.Message)
		})
	}
}
