package middleware

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	appErr "fut-app/internal/errors"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		logger := slog.Default()

		var ve appErr.ValidationErrors
		if errors.As(err, &ve) {
			logger.Warn("validation failed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Any("errors", ve),
			)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"code":    "bad_request",
				"message": "Validation failed",
				"errors":  ve,
			})
			return
		}

		status, httpErr := appErr.ToHTTPError(err)

		logger.Error("request failed",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", status),
			slog.String("error", err.Error()),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(httpErr)
	}
}
