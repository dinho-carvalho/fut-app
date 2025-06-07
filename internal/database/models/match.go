package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"fut-app/internal/database"
)

type StringArray []string

func (a *StringArray) Scan(value interface{}) error {
	// Deserializa do banco para o slice
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal StringArray value: %v", value)
	}
	return json.Unmarshal(bytes, a)
}

func (a StringArray) Value() (driver.Value, error) {
	// Serializa o slice para o banco
	return json.Marshal(a)
}

type Match struct {
	database.Model
	Date     time.Time   `gorm:"not null"`
	Location string      `gorm:"type:varchar(100);not null"`
	TeamA    StringArray `gorm:"type:json"`
	TeamB    StringArray `gorm:"type:json"`
	ScoreA   int         `gorm:"not null"`
	ScoreB   int         `gorm:"not null"`
}
