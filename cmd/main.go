package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"fut-app/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"fut-app/internal/database"
	"fut-app/internal/database/models"
)

func loadEnv() {
	// Load .env files only for non-production environments.
	env := strings.ToLower(strings.TrimSpace(os.Getenv("APP_ENV")))
	if env == "" {
		env = "local"
	}

	switch env {
	case "local", "dev", "development", "stage", "staging", "test":
		_ = godotenv.Load(".env")
		_ = godotenv.Load(".env." + env)
	}
}

func main() {
	loadEnv()
	logger := logger.NewLogger(logger.Config{
		AppName: "fut-app",
	})
	slog.SetDefault(logger)
	db := createDatabase()
	d := InjectDependencies(db, logger)
	r := mux.NewRouter()
	CreateRoutes(r, d)

	slog.Info("🚀 Server is running on port 8080")
	slog.Info("It's time ⚽ ⚽ ⚽ ⚽ ⚽ ⚽")
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("Error starting server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func createDatabase() *database.Database {
	config := database.NewConfig()
	db, err := database.NewDatabase(config)
	if err != nil {
		slog.Error("❌ Failed to connect to the database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("✅ Successfully connected to the database!")

	err = db.AutoMigrate(&models.Player{}, &models.Position{}, &models.Match{}, &models.Rating{})
	if err != nil {
		slog.Error("❌ Error in auto migrate", slog.String("error", err.Error()))
		os.Exit(1)
	}
	return db
}
