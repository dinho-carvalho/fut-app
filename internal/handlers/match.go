package handlers

import (
	"encoding/json"
	"fmt"
	"fut-app/internal/database/models"
	"fut-app/internal/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MatchHandler struct {
	Service services.MatchService
}

func NewMatchHandler(service services.MatchService) *MatchHandler {
	return &MatchHandler{
		Service: service,
	}
}

func (h *MatchHandler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	var match models.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err := h.Service.CreateMatch(match)
	if err != nil {
		http.Error(w, "Failed to create match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MatchHandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches := h.Service.GetAllMatches()
	err := json.NewEncoder(w).Encode(matches)
	if err != nil {
		http.Error(w, "Failed to encode matches", http.StatusInternalServerError)
		return
	}
}

func (h *MatchHandler) GetMatchByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	match, errMatch := h.Service.GetMatchByID(uint(id))
	if errMatch != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	errEncode := json.NewEncoder(w).Encode(match)
	if errEncode != nil {
		http.Error(w, "Failed to encode match", http.StatusInternalServerError)
		return
	}
}

func (h *MatchHandler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var match models.Match

	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err = h.Service.UpdateMatch(match, uint(id)); err != nil {
		http.Error(w, "Failed to update match", http.StatusInternalServerError)
		return
	}
	errEncode := json.NewEncoder(w).Encode(match)
	if errEncode != nil {
		http.Error(w, "Failed to encode match", http.StatusInternalServerError)
		return
	}
}

func (h *MatchHandler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err = h.Service.DeleteMatch(uint(id)); err != nil {
		http.Error(w, "Failed to delete match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("✅ Match deleted successfully")
}
