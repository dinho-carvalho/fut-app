package errors

import (
	"encoding/json"
	"fmt"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []ValidationError

func (ve *ValidationErrors) Error() string {
	out, _ := json.Marshal(ve) // útil para logs
	return fmt.Sprintf("validation failed: %s", string(out))
}

func (ve *ValidationErrors) Append(field, message string) {
	*ve = append(*ve, ValidationError{
		Field:   field,
		Message: message,
	})
}

func (ve *ValidationErrors) HasErrors() bool {
	return len(*ve) > 0
}
