package models

import (
	"time"
)

// Supplier struct describes a single supplier information
type Supplier struct {
	ID                 string
	Name               string
	Description        string
	Status             string
	MainContactID      string
	SecondaryContactID string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type SupplierContact struct {
	ID              string
	Name            string
	Email           string
	Role            string
	HomeNumber      string
	HandphoneNumber string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
