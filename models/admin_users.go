package models

import "time"

// AdminUser struct defines an admin user entity in the application
// Status in the field refers to whether the user is hired/attached etc
// StartDate is the date when the user joined the company
// EndDate is the date when the user left the company
type AdminUser struct {
	User
	Role      Role
	RoleID    string
	Status    string
	StartDate time.Time
	EndDate   time.Time
}
