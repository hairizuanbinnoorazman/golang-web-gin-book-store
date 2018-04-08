package models

import (
	"time"
)

// Role struct descriptive different roles that different users can do on the platform
type Role struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
