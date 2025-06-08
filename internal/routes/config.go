package routes

import (
	"fmt"
	"net/http"

	"fut-app/internal/repositories"
	"fut-app/internal/services"

	"fut-app/internal/handlers"

	"fut-app/internal/database"
	"github.com/gorilla/mux"
)

func CreateRoutes(r *mux.Router, db *database.Database) {
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	players(r, db)
}

func players(r *mux.Router, db *database.Database) {
	repo := repositories.NewPlayer(db.DB)
	service := services.NewPlayerService(repo)
	playerHandler := handlers.PlayerHandler{
		Service: service,
	}
	r.HandleFunc("/players", playerHandler.CreatePlayer).Methods("POST")
	r.HandleFunc("/players", playerHandler.GetPlayers).Methods("GET")
	r.HandleFunc("/players/{id:[0-9]+}", playerHandler.GetPlayerByID).Methods("GET")
	r.HandleFunc("/players/{id:[0-9]+}", playerHandler.UpdatePlayer).Methods("PUT")
	r.HandleFunc("/players/{id:[0-9]+}", playerHandler.DeletePlayer).Methods("DELETE")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
