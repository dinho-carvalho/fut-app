package repositories

import (
	"testing"

	"fut-app/internal/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(&models.Player{}, &models.Match{}, &models.Rating{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
