package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONB map[string]any // change to interface for tests.

func (j *JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert %v to JSONB", value)
	}

	return json.Unmarshal(bytes, j)
}
