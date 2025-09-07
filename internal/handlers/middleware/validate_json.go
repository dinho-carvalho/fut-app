package middleware

import (
	"encoding/json"
	"net/http"
	"sync"

	"fut-app/internal/errors"

	"github.com/go-playground/validator/v10"
)

var (
	validate       = validator.New()
	registerCustom sync.Once
)

func ValidateJSON[T any](next func(http.ResponseWriter, *http.Request, T) error) AppHandler {
	registerCustom.Do(func() {
		if err := validate.RegisterValidation("statslen", func(fl validator.FieldLevel) bool {
			if m, ok := fl.Field().Interface().(map[string]interface{}); ok {
				return len(m) == 6
			}
			return false
		}); err != nil {
			panic(err)
		}
	})

	return func(w http.ResponseWriter, r *http.Request) error {
		var body T
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		if err := dec.Decode(&body); err != nil {
			return errors.ErrInvalidData
		}

		if err := validate.Struct(body); err != nil {
			return errors.ErrInvalidData
		}

		return next(w, r, body)
	}
}
