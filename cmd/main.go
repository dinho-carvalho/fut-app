package main

import (
	"log/slog"
	"net/http"
	"os"

	"fut-app/pkg/logger"

	"github.com/gorilla/mux"

	"fut-app/internal/database"
	"fut-app/internal/database/models"
)

func main() {
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
