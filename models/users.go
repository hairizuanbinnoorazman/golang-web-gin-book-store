package models

import (
	"time"
)

// User struct defines the user entity in the application
type User struct {
	ID                       string    `json:"id"`
	FirstName                string    `json:"first_name"`
	LastName                 string    `json:"last_name"`
	Email                    string    `json:"email"`
	Password                 string    `json:"-"`
	ForgetPasswordToken      string    `json:"-"`
	ForgetPasswordExpiryDate time.Time `json:"-"`
	Activated                bool      `json:"activated"`
	Address                  string    `json:"address"`
	LastLoginAt              time.Time `json:"-"`
	CreatedAt                time.Time `json:"-"`
	UpdatedAt                time.Time `json:"-"`
}

type AdminUser struct {
	User
	Type      string
	RoleID    string
	StartDate time.Time
	EndDate   time.Time
}

func NewUser(firstName, lastName, email, password string) *User {
	return &User{FirstName: firstName, LastName: lastName, Email: email, Password: password}
}

func Update(user *User) {}
