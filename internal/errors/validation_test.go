package errors

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	ve := ValidationError{
		Field:   "name",
		Message: "is required",
	}

	assert.Equal(t, "name", ve.Field)
	assert.Equal(t, "is required", ve.Message)
}

func TestValidationErrors_Error(t *testing.T) {
	tests := []struct {
		name     string
		errors   ValidationErrors
		contains string
	}{
		{
			name:     "single error",
			errors:   ValidationErrors{{Field: "name", Message: "is required"}},
			contains: "validation failed:",
		},
		{
			name: "multiple errors",
			errors: ValidationErrors{
				{Field: "name", Message: "is required"},
				{Field: "email", Message: "is invalid"},
			},
			contains: "validation failed:",
		},
		{
			name:     "empty errors",
			errors:   ValidationErrors{},
			contains: "validation failed:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := tt.errors.Error()
			assert.Contains(t, errMsg, tt.contains)

			// Verify it's valid JSON after "validation failed: "
			prefix := "validation failed: "
			jsonPart := errMsg[len(prefix):]
			var result []ValidationError
			err := json.Unmarshal([]byte(jsonPart), &result)
			assert.NoError(t, err)
			assert.Equal(t, len(tt.errors), len(result))
		})
	}
}

func TestValidationErrors_Append(t *testing.T) {
	var ve ValidationErrors

	assert.False(t, ve.HasErrors())
	assert.Equal(t, 0, len(ve))

	ve.Append("name", "is required")
	assert.True(t, ve.HasErrors())
	assert.Equal(t, 1, len(ve))
	assert.Equal(t, "name", ve[0].Field)
	assert.Equal(t, "is required", ve[0].Message)

	ve.Append("email", "is invalid")
	assert.True(t, ve.HasErrors())
	assert.Equal(t, 2, len(ve))
	assert.Equal(t, "email", ve[1].Field)
	assert.Equal(t, "is invalid", ve[1].Message)
}

func TestValidationErrors_HasErrors(t *testing.T) {
	tests := []struct {
		name     string
		errors   ValidationErrors
		expected bool
	}{
		{"empty", ValidationErrors{}, false},
		{"single error", ValidationErrors{{Field: "name", Message: "required"}}, true},
		{"multiple errors", ValidationErrors{{Field: "name", Message: "required"}, {Field: "email", Message: "invalid"}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.errors.HasErrors())
		})
	}
}

func TestValidationErrors_JSONMarshaling(t *testing.T) {
	ve := ValidationErrors{
		{Field: "name", Message: "is required"},
		{Field: "email", Message: "is invalid"},
	}

	data, err := json.Marshal(ve)
	assert.NoError(t, err)

	var result []ValidationError
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "name", result[0].Field)
	assert.Equal(t, "is required", result[0].Message)
	assert.Equal(t, "email", result[1].Field)
	assert.Equal(t, "is invalid", result[1].Message)
}
