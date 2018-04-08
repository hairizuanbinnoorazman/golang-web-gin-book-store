package models

import (
	"time"
)

// Supplier struct describes a single supplier information
type Supplier struct {
	ID                 string
	Name               string
	Description        string
	Address            string
	Status             string
	MainContactID      string
	SecondaryContactID string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
