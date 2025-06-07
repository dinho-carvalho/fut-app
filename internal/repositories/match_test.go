package repositories

import (
	"testing"

	"fut-app/internal/database"

	"fut-app/internal/database/models"
)

func TestMatchRepository_CreateMatch(t *testing.T) {
	tests := []struct {
		name    string
		match   models.Match
		wantErr bool
	}{
		{
			name: "success",
			match: models.Match{
				Location: "Campo 1",
				TeamA:    []string{"1", "2", "3", "4", "5"},
				TeamB:    []string{"6", "7", "8", "9", "10"},
				ScoreA:   3,
				ScoreB:   2,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupTestDB(t)
			repo := NewMatch(db)

			err := repo.CreateMatch(tt.match)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMatch() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var found models.Match
				result := db.First(&found)
				if result.Error != nil {
					t.Errorf("Failed to find created match: %v", result.Error)
				}
				if found.Location != tt.match.Location {
					t.Errorf("Created match location = %v, want %v", found.Location, tt.match.Location)
				}
			}
		})
	}
}

func TestMatchRepository_GetMatches(t *testing.T) {
	db := setupTestDB(t)
	repo := NewMatch(db)

	// Create test matches
	testMatches := []models.Match{
		{
			Location: "Campo 1",
			TeamA:    []string{"1", "2", "3", "4", "5"},
			TeamB:    []string{"6", "7", "8", "9", "10"},
			ScoreA:   3,
			ScoreB:   2,
		},
		{
			Location: "Campo 2",
			TeamA:    []string{"11", "12", "13", "14", "15"},
			TeamB:    []string{"16", "17", "18", "19", "20"},
			ScoreA:   1,
			ScoreB:   1,
		},
	}
	for _, m := range testMatches {
		if err := db.Create(&m).Error; err != nil {
			t.Fatalf("Failed to create test match: %v", err)
		}
	}

	got := repo.GetMatches()
	if len(got) != len(testMatches) {
		t.Errorf("GetMatches() got = %v matches, want %v", len(got), len(testMatches))
	}
}

func TestMatchRepository_GetMatchByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewMatch(db)

	// Create a test match
	testMatch := models.Match{
		Location: "Campo 1",
		TeamA:    []string{"1", "2", "3", "4", "5"},
		TeamB:    []string{"6", "7", "8", "9", "10"},
		ScoreA:   3,
		ScoreB:   2,
	}
	if err := db.Create(&testMatch).Error; err != nil {
		t.Fatalf("Failed to create test match: %v", err)
	}

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "existing match",
			id:      int(testMatch.ID),
			wantErr: false,
		},
		{
			name:    "non-existing match",
			id:      9999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, err := repo.GetMatchByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && match.Location != testMatch.Location {
				t.Errorf("GetMatchByID() got = %v, want %v", match.Location, testMatch.Location)
			}
		})
	}
}

func TestMatchRepository_UpdateMatch(t *testing.T) {
	db := setupTestDB(t)
	repo := NewMatch(db)

	// Create a test match
	testMatch := models.Match{
		Location: "Campo 1",
		TeamA:    []string{"1", "2", "3", "4", "5"},
		TeamB:    []string{"6", "7", "8", "9", "10"},
		ScoreA:   3,
		ScoreB:   2,
	}
	if err := db.Create(&testMatch).Error; err != nil {
		t.Fatalf("Failed to create test match: %v", err)
	}

	tests := []struct {
		name    string
		match   models.Match
		wantErr bool
	}{
		{
			name: "valid update",
			match: models.Match{
				Model: database.Model{
					ID: testMatch.ID,
				},
				Location: "Campo 2",
				TeamA:    []string{"1", "2", "3", "4", "5"},
				TeamB:    []string{"6", "7", "8", "9", "10"},
				ScoreA:   4,
				ScoreB:   2,
			},
			wantErr: false,
		},
		{
			name: "non-existing match",
			match: models.Match{
				Model: database.Model{
					ID: 9999,
				},
				Location: "Campo 3",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateMatch(tt.match)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateMatch() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var updated models.Match
				if err := db.First(&updated, tt.match.ID).Error; err != nil {
					t.Errorf("Failed to find updated match: %v", err)
				}
				if updated.Location != tt.match.Location {
					t.Errorf("Updated match location = %v, want %v", updated.Location, tt.match.Location)
				}
			}
		})
	}
}

func TestMatchRepository_DeleteMatch(t *testing.T) {
	db := setupTestDB(t)
	repo := NewMatch(db)

	// Create a test match
	testMatch := models.Match{
		Location: "Campo 1",
		TeamA:    []string{"1", "2", "3", "4", "5"},
		TeamB:    []string{"6", "7", "8", "9", "10"},
		ScoreA:   3,
		ScoreB:   2,
	}
	if err := db.Create(&testMatch).Error; err != nil {
		t.Fatalf("Failed to create test match: %v", err)
	}

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "existing match",
			id:      int(testMatch.ID),
			wantErr: false,
		},
		{
			name:    "non-existing match",
			id:      9999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.DeleteMatch(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteMatch() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var found models.Match
				err := db.First(&found, tt.id).Error
				if err == nil {
					t.Error("DeleteMatch() match still exists after deletion")
				}
			}
		})
	}
}
