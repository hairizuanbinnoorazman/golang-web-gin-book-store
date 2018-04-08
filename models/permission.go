package models

import (
	"time"
)

// Permission struct describes the specific action that can be done on the platform
// Permissions can be grouped up to form a role
type Permission struct {
	ID          string
	Name        string
	Description string
	Status      string
	Remarks     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// RolePermissions maps permissions to roles.
// An admin role would be able to overwrite other actions by different people
type RolePermissions struct {
	RoleID       string
	PermissionID string
}
