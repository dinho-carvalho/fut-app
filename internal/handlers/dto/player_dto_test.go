package dto

import (
	"reflect"
	"testing"
)

func TestPlayerDTO_ToDomain(t *testing.T) {
	d := PlayerDTO{
		Name:     "Messi",
		Stats:    map[string]interface{}{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6},
		Position: []string{"RW", "CF"},
	}

	got := d.ToDomain()

	if got.Name != d.Name {
		t.Fatalf("expected name %q, got %q", d.Name, got.Name)
	}
	if !reflect.DeepEqual(got.Stats, d.Stats) {
		t.Fatalf("unexpected stats: %#v", got.Stats)
	}
	if !reflect.DeepEqual(got.Position, d.Position) {
		t.Fatalf("unexpected positions: %#v", got.Position)
	}
}
