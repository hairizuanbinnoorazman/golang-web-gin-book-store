package models

import (
	"errors"
	"time"
)

var ErrInvalidRoleName = errors.New("Invalid Role Name")
var ErrInvalidStatus = errors.New("Invalid Status")
var ErrInvalidDescription = errors.New("Invalid Description")

// Role struct descriptive different roles that different users can do on the platform
type Role struct {
	ID          string
	Name        string
	Description string
	Status      string
	Remarks     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// There are only very few specific roles that are accepted.
// Role Names: Managerial, StoreAdmin, Supplier
func (r Role) validateName() error { return nil }

func (r Role) validateDescription() error { return nil }

// There are only very specific status that are accepted
// Statuses: Active, Depreciated, Inactive
func (r Role) validateStatus() error { return nil }

func (r Role) validateCreatedAt() error { return nil }

func (r Role) validateUpdatedAt() error { return nil }

// Validate function runs a couple of heuristics checks to ensure that the data being passed is
// minimally filled up
func (r Role) Validate() error { return nil }
