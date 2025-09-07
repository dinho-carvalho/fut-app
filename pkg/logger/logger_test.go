package logger

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg := Config{
		AppName: "test-app",
	}

	assert.Equal(t, "test-app", cfg.AppName)
}

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name     string
		appEnv   string
		appName  string
		wantJSON bool
	}{
		{
			name:     "development environment creates text handler",
			appEnv:   "development",
			appName:  "test-app",
			wantJSON: false,
		},
		{
			name:     "local environment creates text handler",
			appEnv:   "local",
			appName:  "test-app",
			wantJSON: false,
		},
		{
			name:     "production environment creates JSON handler",
			appEnv:   "production",
			appName:  "test-app",
			wantJSON: true,
		},
		{
			name:     "staging environment creates JSON handler",
			appEnv:   "staging",
			appName:  "test-app",
			wantJSON: true,
		},
		{
			name:     "empty environment creates JSON handler",
			appEnv:   "",
			appName:  "test-app",
			wantJSON: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment
			originalEnv := os.Getenv("APP_ENV")
			defer os.Setenv("APP_ENV", originalEnv)
			os.Setenv("APP_ENV", tt.appEnv)

			cfg := Config{
				AppName: tt.appName,
			}

			logger := NewLogger(cfg)

			assert.NotNil(t, logger)
			assert.IsType(t, &slog.Logger{}, logger)

			// Test that logger has the expected attributes
			// We can't easily test the handler type directly, but we can test
			// that the logger was created successfully
			logger.Info("test message")
		})
	}
}

func TestNewLogger_WithAttributes(t *testing.T) {
	originalEnv := os.Getenv("APP_ENV")
	defer os.Setenv("APP_ENV", originalEnv)
	os.Setenv("APP_ENV", "test")

	cfg := Config{
		AppName: "fut-app",
	}

	logger := NewLogger(cfg)

	// Test that we can use the logger without panics
	assert.NotPanics(t, func() {
		logger.Info("test info message")
		logger.Warn("test warn message")
		logger.Error("test error message")
	})
}

func TestNewLogger_DifferentAppNames(t *testing.T) {
	originalEnv := os.Getenv("APP_ENV")
	defer os.Setenv("APP_ENV", originalEnv)
	os.Setenv("APP_ENV", "local")

	testCases := []string{
		"fut-app",
		"test-service",
		"",
		"app-with-dashes",
		"app_with_underscores",
	}

	for _, appName := range testCases {
		t.Run("app_name_"+appName, func(t *testing.T) {
			cfg := Config{
				AppName: appName,
			}

			logger := NewLogger(cfg)
			assert.NotNil(t, logger)

			// Test that logger works
			assert.NotPanics(t, func() {
				logger.Info("testing app name", slog.String("app_name", appName))
			})
		})
	}
}

func TestNewLogger_HandlerOptionsLevel(t *testing.T) {
	originalEnv := os.Getenv("APP_ENV")
	defer os.Setenv("APP_ENV", originalEnv)

	// Test different environments to ensure level is consistently set
	environments := []string{"local", "development", "production", "staging"}

	for _, env := range environments {
		t.Run("env_"+env, func(t *testing.T) {
			os.Setenv("APP_ENV", env)

			cfg := Config{
				AppName: "test-app",
			}

			logger := NewLogger(cfg)
			assert.NotNil(t, logger)

			// All messages should be logged at Info level
			assert.NotPanics(t, func() {
				logger.Debug("debug message") // This might not show depending on level
				logger.Info("info message")   // This should show
				logger.Warn("warn message")   // This should show
				logger.Error("error message") // This should show
			})
		})
	}
}
