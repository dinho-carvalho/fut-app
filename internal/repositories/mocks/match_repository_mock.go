package mocks

import (
	"fut-app/internal/database/models"
)

type MatchRepositoryMock struct {
	CreateMatchFunc  func(models.Match) error
	GetMatchesFunc   func() []models.Match
	GetMatchByIDFunc func(int) (models.Match, error)
	UpdateMatchFunc  func(models.Match) error
	DeleteMatchFunc  func(int) error
}

func (m *MatchRepositoryMock) CreateMatch(match models.Match) error {
	if m.CreateMatchFunc != nil {
		return m.CreateMatchFunc(match)
	}
	return nil
}

func (m *MatchRepositoryMock) GetMatches() []models.Match {
	if m.GetMatchesFunc != nil {
		return m.GetMatchesFunc()
	}
	return []models.Match{}
}

func (m *MatchRepositoryMock) GetMatchByID(id int) (models.Match, error) {
	if m.GetMatchByIDFunc != nil {
		return m.GetMatchByIDFunc(id)
	}
	return models.Match{}, nil
}

func (m *MatchRepositoryMock) UpdateMatch(match models.Match) error {
	if m.UpdateMatchFunc != nil {
		return m.UpdateMatchFunc(match)
	}
	return nil
}

func (m *MatchRepositoryMock) DeleteMatch(id int) error {
	if m.DeleteMatchFunc != nil {
		return m.DeleteMatchFunc(id)
	}
	return nil
}
