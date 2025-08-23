package services

import (
	"errors"
	"fut-app/internal/database/models"
	"fut-app/internal/repositories"
)

type (
	MatchService interface {
		CreateMatch(models.Match) error
		GetAllMatches() []models.Match
		GetMatchByID(uint) (*models.Match, error)
		UpdateMatch(models.Match, uint) error
		DeleteMatch(uint) error
	}
	matchimpl struct {
		repo repositories.MatchRepository
	}
)

func NewMatchService(repo repositories.MatchRepository) MatchService {
	return &matchimpl{repo: repo}
}

func (s *matchimpl) CreateMatch(match models.Match) error {
	return s.repo.CreateMatch(match)
}

func (s *matchimpl) GetAllMatches() []models.Match {
	return s.repo.GetMatches()
}

func (s *matchimpl) GetMatchByID(id uint) (*models.Match, error) {
	return s.repo.GetMatchByID(id)
}

func (s *matchimpl) UpdateMatch(match models.Match, id uint) error {
	p, err := s.repo.GetMatchByID(id)
	if err != nil {
		return errors.New("error getting match") // TODO: refactor error handling
	}
	if p == nil {
		return errors.New("match not found") // TODO: refactor error handling
	}
	return s.repo.UpdateMatch(match)
}

func (s *matchimpl) DeleteMatch(id uint) error {
	return s.repo.DeleteMatch(id)
}
