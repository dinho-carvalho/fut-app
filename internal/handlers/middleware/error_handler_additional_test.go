package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	appErrors "fut-app/internal/errors"
)

func TestAppHandler_ServesHTTP_NoError(t *testing.T) {
	h := AppHandler(func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("success"))
		return nil
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "success", rr.Body.String())
}

func TestAppHandler_ServesHTTP_MultipleValidationErrors(t *testing.T) {
	h := AppHandler(func(w http.ResponseWriter, r *http.Request) error {
		var ve appErrors.ValidationErrors
		ve.Append("name", "is required")
		ve.Append("email", "is invalid")
		ve.Append("age", "must be positive")
		return &ve
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/players", nil)
	h.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "bad_request", response["code"])
	assert.Equal(t, "Validation failed", response["message"])

	errors, ok := response["errors"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, errors, 3)
}

func TestAppHandler_ServesHTTP_DifferentHTTPMethods(t *testing.T) {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			h := AppHandler(func(w http.ResponseWriter, r *http.Request) error {
				return appErrors.ErrNotFound
			})

			rr := httptest.NewRecorder()
			req := httptest.NewRequest(method, "/test", nil)
			h.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusNotFound, rr.Code)
		})
	}
}

func TestAppHandler_ServesHTTP_EmptyValidationErrors(t *testing.T) {
	h := AppHandler(func(w http.ResponseWriter, r *http.Request) error {
		var ve appErrors.ValidationErrors
		// Empty validation errors should still be handled
		return &ve
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/players", nil)
	h.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "bad_request", response["code"])
	assert.Equal(t, "Validation failed", response["message"])
}
