package models

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email string `json:"email"`
	// Expire     time.Time `json:"expire"`
	// Created_at time.Time `json:"created_at"`
	// Updated_at time.Time `json:"updated_at"`
	// Teraffic   int       `json:"teraffic"`
	// Uuid       uuid.UUID `json:"uuid"`
	Active bool `json:"active"`
}

// This method simply returns the JSON-encoded representation of the users struct.
func (u Users) Value() ([]byte, error) {
	return json.Marshal(u)
}

// This method simply decodes a JSON-encoded value into the struct fields.
func (u *Users) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &u)
}
