package models

import (
	"errors"
	"time"
)

var ErrInvalidRoleName = errors.New("Invalid Role Name")
var ErrInvalidRoleStatus = errors.New("Invalid Status")
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
func (r Role) validateName() error {
	if r.Name == "manager" || r.Name == "storeadmin" || r.Name == "supplier" {
		return nil
	}
	return ErrInvalidRoleName
}

func (r Role) validateDescription() error {
	return nil
}

// There are only very specific status that are accepted
// Statuses: Active, Depreciated, Inactive
// Status of the role has to be lowercased
func (r Role) validateStatus() error {
	if r.Status == "active" || r.Status == "deactivated" || r.Status == "depreciated" {
		return nil
	}
	return ErrInvalidRoleStatus
}

// Validate function runs a couple of heuristics checks to ensure that the data being passed is
// minimally filled up
func (r Role) Validate() error { return nil }

// NewRole functions creates a new role function with predefined values such as ID etc to be used
func NewRole(name, description string) (*Role, error) {
	return &Role{}, nil
}
