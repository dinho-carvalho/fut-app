package domain

type (
	Player struct {
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
