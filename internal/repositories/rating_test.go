package repositories

import (
	"testing"

	"fut-app/internal/database"

	"fut-app/internal/database/models"
)

func TestRatingRepository_CreateRating(t *testing.T) {
	tests := []struct {
		name    string
		rating  models.Rating
		wantErr bool
	}{
		{
			name: "success",
			rating: models.Rating{
				MatchID:       1,
				PlayerID:      1,
				RatedPlayerID: 2,
				Finishing:     80,
				Passing:       75,
				Speed:         85,
				Defense:       70,
				Stamina:       90,
				Highlight:     95,
			},
			wantErr: false,
		},
		{
			name: "invalid rating values",
			rating: models.Rating{
				MatchID:       1,
				PlayerID:      1,
				RatedPlayerID: 2,
				Finishing:     101, // Invalid value > 100
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupTestDB(t)
			repo := NewRating(db)

			err := repo.CreateRating(tt.rating)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRating() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var found models.Rating
				result := db.First(&found)
				if result.Error != nil {
					t.Errorf("Failed to find created rating: %v", result.Error)
				}
				if found.Finishing != tt.rating.Finishing {
					t.Errorf("Created rating finishing = %v, want %v", found.Finishing, tt.rating.Finishing)
				}
			}
		})
	}
}

func TestRatingRepository_GetRatings(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRating(db)

	// Create test ratings
	testRatings := []models.Rating{
		{
			MatchID:       1,
			PlayerID:      1,
			RatedPlayerID: 2,
			Finishing:     80,
		},
		{
			MatchID:       1,
			PlayerID:      2,
			RatedPlayerID: 1,
			Finishing:     85,
		},
	}
	for _, r := range testRatings {
		if err := db.Create(&r).Error; err != nil {
			t.Fatalf("Failed to create test rating: %v", err)
		}
	}

	got := repo.GetRatings()
	if len(got) != len(testRatings) {
		t.Errorf("GetRatings() got = %v ratings, want %v", len(got), len(testRatings))
	}
}

func TestRatingRepository_GetRatingByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRating(db)

	// Create a test rating
	testRating := models.Rating{
		MatchID:       1,
		PlayerID:      1,
		RatedPlayerID: 2,
		Finishing:     80,
	}
	if err := db.Create(&testRating).Error; err != nil {
		t.Fatalf("Failed to create test rating: %v", err)
	}

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "existing rating",
			id:      int(testRating.ID),
			wantErr: false,
		},
		{
			name:    "non-existing rating",
			id:      9999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rating, err := repo.GetRatingByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRatingByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && rating.Finishing != testRating.Finishing {
				t.Errorf("GetRatingByID() got = %v, want %v", rating.Finishing, testRating.Finishing)
			}
		})
	}
}

func TestRatingRepository_UpdateRating(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRating(db)

	// Create a test rating
	testRating := models.Rating{
		MatchID:       1,
		PlayerID:      1,
		RatedPlayerID: 2,
		Finishing:     80,
	}
	if err := db.Create(&testRating).Error; err != nil {
		t.Fatalf("Failed to create test rating: %v", err)
	}

	tests := []struct {
		name    string
		rating  models.Rating
		wantErr bool
	}{
		{
			name: "valid update",
			rating: models.Rating{
				Model: database.Model{
					ID: testRating.ID,
				},
				MatchID:   1,
				PlayerID:  1,
				Finishing: 90,
			},
			wantErr: false,
		},
		{
			name: "non-existing rating",
			rating: models.Rating{
				Model: database.Model{
					ID: 9999,
				},
				Finishing: 85,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateRating(tt.rating)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRating() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var updated models.Rating
				if err := db.First(&updated, tt.rating.ID).Error; err != nil {
					t.Errorf("Failed to find updated rating: %v", err)
				}
				if updated.Finishing != tt.rating.Finishing {
					t.Errorf("Updated rating finishing = %v, want %v", updated.Finishing, tt.rating.Finishing)
				}
			}
		})
	}
}

func TestRatingRepository_DeleteRating(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRating(db)

	// Create a test rating
	testRating := models.Rating{
		MatchID:       1,
		PlayerID:      1,
		RatedPlayerID: 2,
		Finishing:     80,
	}
	if err := db.Create(&testRating).Error; err != nil {
		t.Fatalf("Failed to create test rating: %v", err)
	}

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "existing rating",
			id:      int(testRating.ID),
			wantErr: false,
		},
		{
			name:    "non-existing rating",
			id:      9999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.DeleteRating(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRating() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				var found models.Rating
				err := db.First(&found, tt.id).Error
				if err == nil {
					t.Error("DeleteRating() rating still exists after deletion")
				}
			}
		})
	}
}
