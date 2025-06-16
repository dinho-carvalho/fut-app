package routes

import (
	"fmt"
	"log/slog"
	"net/http"

	"fut-app/internal/handlers/dto"
	"fut-app/internal/handlers/middleware"

	"fut-app/internal/repositories"
	"fut-app/internal/services"

	"fut-app/internal/handlers"

	"fut-app/internal/database"

	"github.com/gorilla/mux"
)

func CreateRoutes(r *mux.Router, db *database.Database, logger *slog.Logger) { // TODO criar app dependency e remover repositories daqui.
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	players(r, db, logger)
}

func players(r *mux.Router, db *database.Database, logger *slog.Logger) {
	repo := repositories.NewPlayer(db.DB, logger)
	service := services.NewPlayerService(repo)
	playerHandler := handlers.NewPlayerHandler(service)
<<<<<<< HEAD
	r.HandleFunc("/players", playerHandler.CreatePlayer).Methods("POST")
=======
	r.HandleFunc("/players", middleware.ValidateJSON[dto.PlayerDTO](playerHandler.CreatePlayer)).Methods("POST")
>>>>>>> 0f4fc5e (feat: create player)
	r.HandleFunc("/players", playerHandler.GetPlayers).Methods("GET")
	r.HandleFunc("/players/{id:[0-9]+}", playerHandler.GetPlayerByID).Methods("GET")
	r.HandleFunc("/players/{id:[0-9]+}", playerHandler.UpdatePlayer).Methods("PUT")
	r.HandleFunc("/players/{id:[0-9]+}", playerHandler.DeletePlayer).Methods("DELETE")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
