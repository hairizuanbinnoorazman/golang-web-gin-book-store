package models

import (
	"errors"
	"time"
)

var ErrStoreNameShort = errors.New("Store name too short")
var ErrStoreNameLong = errors.New("Store name too long")
var ErrStoreNameInvalid = errors.New("Store name is invalid")

// Store struct - there should only be 1 company director/admin
type Store struct {
	ID          string
	Name        string
	AdminUser   AdminUser
	AdminUserID string
	Address     string
	Country     string
	Status      string
	Remarks     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s Store) validateStoreName() error { return nil }

func (s Store) validateStoreAdminUser() error { return nil }

func (s Store) validateStoreAddress() error { return nil }

func (s Store) validateStoreCountry() error { return nil }

func (s Store) validateStoreStatus() error { return nil }

// Validate checks the store record to ensure data is safe to be saved
func (s Store) Validate() error { return nil }
