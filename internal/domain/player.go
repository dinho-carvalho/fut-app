package domain

import "fut-app/internal/errors"

type (
	Player struct {
		ID       uint
		Name     string
		Stats    map[string]interface{}
		Position []string
	}

	Position struct {
		Name string
	}
)

func NewPlayer(name string, stats map[string]interface{}, position []string) *Player {
	return &Player{
		Name:     name,
		Stats:    stats,
		Position: position,
	}
}

func (p Player) Validate() error {
	var errs errors.ValidationErrors

	if p.Name == "" {
		errs.Append("name", "Name is required")
	}
	if len(p.Stats) != 6 {
		errs.Append("stats", "Stats must contain exactly 6 keys")
	}
	if len(p.Position) == 0 {
		errs.Append("positions", "At least one position is required")
	}

	if errs.HasErrors() {
		return &errs
	}
	return nil
}
