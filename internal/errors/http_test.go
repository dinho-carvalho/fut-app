package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPError(t *testing.T) {
	httpErr := HTTPError{
		Code:    "test_error",
		Message: "Test error message",
	}

	assert.Equal(t, "test_error", httpErr.Code)
	assert.Equal(t, "Test error message", httpErr.Message)
}

func TestToHTTPError(t *testing.T) {
	tests := []struct {
		name           string
		inputError     error
		expectedStatus int
		expectedCode   string
		expectedMsg    string
	}{
		{
			name:           "ErrNotFound",
			inputError:     ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedCode:   "not_found",
			expectedMsg:    "Resource not found",
		},
		{
			name:           "ErrBadRequest",
			inputError:     ErrBadRequest,
			expectedStatus: http.StatusBadRequest,
			expectedCode:   "bad_request",
			expectedMsg:    "Bad request",
		},
		{
			name:           "ErrAlreadyExists",
			inputError:     ErrAlreadyExists,
			expectedStatus: http.StatusConflict,
			expectedCode:   "already_exists",
			expectedMsg:    "Resource already exists",
		},
		{
			name:           "ErrInvalidData",
			inputError:     ErrInvalidData,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedCode:   "invalid_data",
			expectedMsg:    "Invalid data provided",
		},
		{
			name:           "ErrUnauthorized",
			inputError:     ErrUnauthorized,
			expectedStatus: http.StatusUnauthorized,
			expectedCode:   "unauthorized",
			expectedMsg:    "Unauthorized",
		},
		{
			name:           "ErrForbidden",
			inputError:     ErrForbidden,
			expectedStatus: http.StatusForbidden,
			expectedCode:   "forbidden",
			expectedMsg:    "Forbidden",
		},
		{
			name:           "ErrDatabase",
			inputError:     ErrDatabase,
			expectedStatus: http.StatusInternalServerError,
			expectedCode:   "database_error",
			expectedMsg:    "Database error",
		},
		{
			name:           "Unknown error",
			inputError:     errors.New("some unknown error"),
			expectedStatus: http.StatusInternalServerError,
			expectedCode:   "internal_error",
			expectedMsg:    "Unexpected error",
		},
		{
			name:           "Nil error defaults to internal",
			inputError:     nil,
			expectedStatus: http.StatusInternalServerError,
			expectedCode:   "internal_error",
			expectedMsg:    "Unexpected error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, httpErr := ToHTTPError(tt.inputError)

			assert.Equal(t, tt.expectedStatus, status)
			assert.Equal(t, tt.expectedCode, httpErr.Code)
			assert.Equal(t, tt.expectedMsg, httpErr.Message)
		})
	}
}

func TestToHTTPError_WithWrappedErrors(t *testing.T) {
	// Test wrapped errors to ensure errors.Is works correctly
	wrappedNotFound := errors.Join(ErrNotFound, errors.New("additional context"))

	status, httpErr := ToHTTPError(wrappedNotFound)

	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, "not_found", httpErr.Code)
	assert.Equal(t, "Resource not found", httpErr.Message)
}
