package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorConstants(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{"ErrNotFound", ErrNotFound, "resource not found"},
		{"ErrAlreadyExists", ErrAlreadyExists, "resource already exists"},
		{"ErrInvalidData", ErrInvalidData, "invalid data provided"},
		{"ErrBadRequest", ErrBadRequest, "bad request"},
		{"ErrInternal", ErrInternal, "internal server error"},
		{"ErrDatabase", ErrDatabase, "database error"},
		{"ErrUnauthorized", ErrUnauthorized, "unauthorized"},
		{"ErrForbidden", ErrForbidden, "forbidden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

func TestErrorsAreDistinct(t *testing.T) {
	errors := []error{
		ErrNotFound,
		ErrAlreadyExists,
		ErrInvalidData,
		ErrBadRequest,
		ErrInternal,
		ErrDatabase,
		ErrUnauthorized,
		ErrForbidden,
	}

	for i, err1 := range errors {
		for j, err2 := range errors {
			if i != j {
				assert.NotEqual(t, err1, err2, "Errors should be distinct")
			}
		}
	}
}
