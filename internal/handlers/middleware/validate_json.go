package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate       = validator.New()
	registerCustom sync.Once
)

type ValidationErrorResponse struct {
	Code    string               `json:"code"`
	Message string               `json:"message"`
	Errors  []FieldErrorResponse `json:"errors"`
}

type FieldErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateJSON[T any](next func(http.ResponseWriter, *http.Request, T)) http.HandlerFunc {
	logger := slog.Default()
	registerCustom.Do(func() {
		_ = validate.RegisterValidation("statslen", func(fl validator.FieldLevel) bool {
			if m, ok := fl.Field().Interface().(map[string]interface{}); ok {
				return len(m) == 6
			}
			return false
		})
	})

	return func(w http.ResponseWriter, r *http.Request) {
		var body T

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			logger.Error("failed to decode request body", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ValidationErrorResponse{
				Code:    "bad_request",
				Message: "Invalid JSON body",
				Errors:  []FieldErrorResponse{},
			})
			return
		}

		if err := validate.Struct(body); err != nil {
			var errors []FieldErrorResponse
			for _, e := range err.(validator.ValidationErrors) {
				field := strings.ToLower(e.Field())
				message := validationMessage(e)
				errors = append(errors, FieldErrorResponse{Field: field, Message: message})
			}

			logger.Warn("validation failed",
				slog.String("path", r.URL.Path),
				slog.String("method", r.Method),
				slog.Any("errors", errors),
			)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ValidationErrorResponse{
				Code:    "bad_request",
				Message: "Validation failed",
				Errors:  errors,
			})
			return
		}

		next(w, r, body)
	}
}

func validationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Must contain at least one item"
	case "statslen":
		return "Stats must contain exactly 6 keys"
	default:
		return "Invalid value"
	}
}
