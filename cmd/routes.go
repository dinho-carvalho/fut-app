package main

import (
	"fmt"
	"net/http"

	"fut-app/internal/handlers/dto"
	"fut-app/internal/handlers/middleware"

	"fut-app/internal/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(r *mux.Router, d Dependencies) { // TODO criar app dependency e remover repositories daqui.
	r.HandleFunc("/health", HealthCheckHandler).Methods(http.MethodGet)
	players(r, d)
}

func players(r *mux.Router, d Dependencies) {
	playerHandler := handlers.NewPlayerHandler(d.RegisterPlayerUseCase)

	r.Handle("/players",
		middleware.ValidateJSON[dto.PlayerDTO](playerHandler.CreatePlayer),
	).Methods(http.MethodPost)

	//r.Handle("/players/{id}", middleware.AppHandler(playerHandler.GetPlayerByID)).Methods(http.MethodGet)
	//
	//r.Handle("/players", middleware.AppHandler(playerHandler.GetPlayers)).Methods(http.MethodGet)
	//
	//r.Handle("/players/{id:[0-9]+}",
	//	middleware.ValidateJSON[dto.PlayerDTO](playerHandler.UpdatePlayer),
	//).Methods(http.MethodPut)
	//
	//r.Handle("/players/{id:[0-9]+}", middleware.AppHandler(playerHandler.DeletePlayer)).Methods(http.MethodDelete)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
