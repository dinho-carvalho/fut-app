package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fut-app/internal/database/models"
	"fut-app/internal/handlers/dto"
	"fut-app/internal/services"

	"github.com/gorilla/mux"
)

type PlayerHandler struct {
	Service services.PlayerService
}

func NewPlayerHandler(service services.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		Service: service,
	}
}

func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request, p dto.PlayerDTO) {
	err := h.Service.CreatePlayer(p.ToDomain())
	if err != nil {
		http.Error(w, "Failed to create player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players := h.Service.GetAllPlayers()
	err := json.NewEncoder(w).Encode(players)
	if err != nil {
		http.Error(w, "Failed to encode players", http.StatusInternalServerError)
		return
	}
}

func (h *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	player, errPlayer := h.Service.GetPlayerByID(uint(id))
	if errPlayer != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	errEncode := json.NewEncoder(w).Encode(player)
	if errEncode != nil {
		http.Error(w, "Failed to encode player", http.StatusInternalServerError)
		return
	}
}

func (h *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var player models.Player

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err = h.Service.UpdatePlayer(player, uint(id)); err != nil {
		http.Error(w, "Failed to update player", http.StatusInternalServerError)
		return
	}
	errEncode := json.NewEncoder(w).Encode(player)
	if errEncode != nil {
		http.Error(w, "Failed to encode player", http.StatusInternalServerError)
		return
	}
}

func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err = h.Service.DeletePlayer(uint(id)); err != nil {
		http.Error(w, "Failed to delete player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("✅ Player deleted successfully")
}
