package services

import (
	"fut-app/internal/database/models"
	"fut-app/internal/repositories"
)

type MatchService struct {
	repo repositories.MatchRepository
}

func NewMatchService(repo repositories.MatchRepository) *MatchService {
	return &MatchService{repo: repo}
}

func (s *MatchService) CreateMatch(match models.Match) error {
	return s.repo.CreateMatch(match)
}

func (s *MatchService) GetAllMatches() []models.Match {
	return s.repo.GetMatches()
}

func (s *MatchService) GetMatchByID(id int) (models.Match, error) {
	return s.repo.GetMatchByID(id)
}

func (s *MatchService) UpdateMatch(match models.Match) error {
	return s.repo.UpdateMatch(match)
}

func (s *MatchService) DeleteMatch(id int) error {
	return s.repo.DeleteMatch(id)
}
