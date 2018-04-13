package models

import (
	"time"
)

// User struct defines the user entity in the application
type User struct {
	ID                       string `json:"id"`
	FirstName                string `json:"first_name"`
	LastName                 string `json:"last_name"`
	Email                    string `json:"email"`
	Address                  string `json:"address"`
	password                 string
	forgetPasswordToken      string
	forgetPasswordExpiryDate time.Time
	activationToken          string
	activationExpiryDate     time.Time
	activated                bool
	lastLoginAt              time.Time
	createdAt                time.Time
	updatedAt                time.Time
}

type AdminUser struct {
	User
	RoleID    string
	Status    string
	StartDate time.Time
	EndDate   time.Time
}

// NewUser returns a new user instance. During instantiation, several fields will need to be
// instantiated; e.g. ID, Activated, CreatedAt, UpdatedAt, etc
func NewUser(firstName, lastName, email, password string) *User {
	return &User{FirstName: firstName, LastName: lastName, Email: email, password: password}
}

func (u User) validateName() error { return nil }

func (u User) validatePassword() error { return nil }

func (u User) validateEmail() error { return nil }

func (u User) validateAddress() error { return nil }

func (u *User) setPassword(password string) error { return nil }

// Validate checks all entries are correct before passing itself to a service to be saved
func (u User) Validate() error { return nil }

// ForgetPassword resets the forget password token to a random UUID as well as resets the
// forget password expiry token. The function will return the forgetPasswordToken
func (u *User) ForgetPassword() (string, error) { return "", nil }

// ChangePasswordFromForget requires you to provide the forget password token. This function
// will then check the forgetPasswordToken if its correct and alters it accordingly
func (u *User) ChangePasswordFromForget(forgetPasswordToken string) error { return nil }

// ReactivateToken resets the activation token in the case the user did not activate the account
// in time. Returns an activationToken
func (u *User) ReactivateToken() (string, error) { return "", nil }

// Activate user. You would need to provide a activation token to check if it correct.
// If correct, it would return the status of the user which should be true or false
func (u *User) Activate(activationToken string) (bool, error) { return false, nil }

// IsActive checks if the user has already activated the token or not
func (u *User) IsActive() bool { return false }
