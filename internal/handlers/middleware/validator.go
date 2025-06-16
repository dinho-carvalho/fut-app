package middleware

import "github.com/go-playground/validator/v10"

func statsLength(fl validator.FieldLevel) bool {
	stats, ok := fl.Field().Interface().(map[string]interface{})
	if !ok {
		return false
	}
	return len(stats) == 6
}
