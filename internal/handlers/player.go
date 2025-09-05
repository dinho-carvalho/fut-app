package handlers

import (
	"net/http"

	"fut-app/internal/usecase"

	"fut-app/internal/handlers/httprespond"

	"fut-app/internal/handlers/dto"
)

type PlayerHandler struct {
	usecase.RegisterPlayerUseCase
}

func NewPlayerHandler(p usecase.RegisterPlayerUseCase) *PlayerHandler {
	return &PlayerHandler{
		RegisterPlayerUseCase: p,
	}
}

func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request, p dto.PlayerDTO) error {
	newPlayer, err := h.RegisterPlayerUseCase.Execute(p.ToDomain())
	if err != nil {
		return err
	}
	return httprespond.JSON(w, http.StatusCreated, newPlayer)
}

//func (h *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) error {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		return errors.ErrBadRequest
//	}
//
//	player, err := h.Service.GetPlayerByID(uint(id))
//	if err != nil {
//		return err
//	}
//
//	return httprespond.JSON(w, http.StatusOK, player)
//}

//func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) error {
//	players := h.Service.GetAllPlayers()
//	return httprespond.JSON(w, http.StatusOK, players)
//}

//func (h *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request, p dto.PlayerDTO) error {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		return errors.ErrBadRequest
//	}
//	var player models.Player
//
//	if err = h.Service.UpdatePlayer(p.ToDomain(), uint(id)); err != nil {
//		return err
//	}
//	return httprespond.JSON(w, http.StatusOK, player)
//}

//func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) error {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		return errors.ErrBadRequest
//	}
//
//	if err = h.Service.DeletePlayer(uint(id)); err != nil {
//		return err
//	}
//	return httprespond.JSON(w, http.StatusNoContent, nil)
//}
