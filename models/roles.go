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
// Role names need to be lowercased
func (r Role) validateName() error { return nil }

func (r Role) validateDescription() error { return nil }

// There are only very specific status that are accepted
// Statuses: Active, Depreciated, Inactive
// Status of the role has to be lowercased
func (r Role) validateStatus() error { return nil }

// Need to check the value is actually created on initial save
func (r Role) validateCreatedAt() error { return nil }

// Need to check the value is actuall created on initial save and subsequent saves
func (r Role) validateUpdatedAt() error { return nil }

// Validate function runs a couple of heuristics checks to ensure that the data being passed is
// minimally filled up
func (r Role) Validate() error { return nil }
