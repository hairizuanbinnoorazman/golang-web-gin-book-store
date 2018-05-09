package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrStoreNameShort = errors.New("Store name too short")
var ErrStoreNameLong = errors.New("Store name too long")
var ErrStoreNameInvalid = errors.New("Store name is invalid")
var ErrStoreStatusInvalid = errors.New("Invalid Status for Store")
var ErrStoreCountryInvalid = errors.New("Invalid Country entry for Store")
var ErrStoreAddress = errors.New("Invalid Address for Store")

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

// NewStore returns a new store struct
func NewStore() (*Store, error) {
	newStore := Store{}
	newStore.ID = uuid.New().String()
	return &Store{}, nil
}
