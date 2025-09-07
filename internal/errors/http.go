package errors

import (
	"errors"
	"net/http"
)

type HTTPError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ToHTTPError(err error) (int, HTTPError) {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound, HTTPError{"not_found", "Resource not found"}
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest, HTTPError{"bad_request", "Bad request"}
	case errors.Is(err, ErrAlreadyExists):
		return http.StatusConflict, HTTPError{"already_exists", "Resource already exists"}
	case errors.Is(err, ErrInvalidData):
		return http.StatusUnprocessableEntity, HTTPError{"invalid_data", "Invalid data provided"}
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized, HTTPError{"unauthorized", "Unauthorized"}
	case errors.Is(err, ErrForbidden):
		return http.StatusForbidden, HTTPError{"forbidden", "Forbidden"}
	case errors.Is(err, ErrDatabase):
		return http.StatusInternalServerError, HTTPError{"database_error", "Database error"}
	default:
		return http.StatusInternalServerError, HTTPError{"internal_error", "Unexpected error"}
	}
}
