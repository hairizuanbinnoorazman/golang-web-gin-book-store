package models

import (
	"time"
)

// ACL maps permissions to roles.
// An admin role would be able to overwrite other actions by different people
type ACL struct {
	RoleID      string
	AdminUserID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
