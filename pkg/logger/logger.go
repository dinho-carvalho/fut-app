package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	AppName string
}

func NewLogger(cfg Config) *slog.Logger {
	env := os.Getenv("APP_ENV")
	level := slog.LevelInfo

	var handler slog.Handler
	if env == "local" || env == "development" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}

	return slog.New(handler).With(
		slog.String("app", cfg.AppName),
		slog.String("env", env),
	)
}
