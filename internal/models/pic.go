package models

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Pic struct {
	gorm.Model
	Source      string
	Link        string
	Srcs        Srcs
	Title       string
	Description string
}

type Srcs []string

func (s Srcs) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Srcs) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), s)
}
