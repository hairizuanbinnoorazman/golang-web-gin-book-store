package models

import (
	"time"
)

// Store struct - there should only be 1 company director/admin
type Store struct {
	ID            string
	Name          string
	CompanyLeadID string // AdminUserID
	Address       string
	Country       string
	Status        string
	Remarks       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (s Store) validateName() error { return nil }

func (s Store) validateCompanyLead() error { return nil }

func (s Store) validateAddress() error { return nil }

func (s Store) validateCountry() error { return nil }

func (s Store) validateStatus() error { return nil }

// Validate checks the store record to ensure data is safe to be saved
func (s Store) Validate() error { return nil }
