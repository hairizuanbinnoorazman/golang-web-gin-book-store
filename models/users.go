package models

import (
	"errors"
	"regexp"
	"time"
)

// ErrNameShort is an error that is to be raised if First Name or Last Name provided is too short
var ErrNameShort = errors.New("Name cannot be shorter than 4 characters")
var ErrNameLong = errors.New("Name cannot be longer than 120 characters")
var ErrAddressLong = errors.New("Address cannot be longer than 120 characters")
var ErrPasswordShort = errors.New("Password cannot be shorter than 8 characters")
var ErrPasswordLong = errors.New("Password cannot be longer than 120 characters")
var ErrPasswordInvalid = errors.New("Password requires at least 1 capital letter, 1 small letter and a number")
var ErrEmailInvalid = errors.New("Email is invalid")
var ErrActivationTokenInvalid = errors.New("Activation Token is invalid")
var ErrForgetPasswordTokenInvalid = errors.New("Forget Password Token is invalid")

// User struct defines the user entity in the application
// Fields need to be set as public in order for other packages to access them
type User struct {
	ID                       string    `json:"id"`
	FirstName                string    `json:"first_name"`
	LastName                 string    `json:"last_name"`
	Email                    string    `json:"email"`
	Address                  string    `json:"address"`
	Password                 string    `json:"-"`
	ForgetPasswordToken      string    `json:"-"`
	ForgetPasswordExpiryDate time.Time `json:"-"`
	ActivationToken          string    `json:"-"`
	ActivationExpiryDate     time.Time `json:"-"`
	Activated                bool      `json:"-"`
	LastLoginAt              time.Time `json:"-"`
	CreatedAt                time.Time `json:"-"`
	UpdatedAt                time.Time `json:"-"`
}

// AdminUser struct defines an admin user entity in the application
// Status in the field refers to whether the user is hired/attached etc
// StartDate is the date when the user joined the company
// EndDate is the date when the user left the company
type AdminUser struct {
	User
	RoleID    string
	Status    string
	StartDate time.Time
	EndDate   time.Time
}

// NewUser returns a new user instance. During instantiation, several fields will need to be
// instantiated; e.g. ID, Activated, CreatedAt, UpdatedAt, etc
// NewUser function would be able to return its own set of errors as it would run a validate function before
// returning a reference of the user back
func NewUser(firstName, lastName, email, password string) (*User, error) {
	return &User{FirstName: firstName, LastName: lastName, Email: email, Password: password}, nil
}

func (u User) validateName() error {
	if len(u.FirstName) < 4 || len(u.LastName) < 4 {
		return ErrNameShort
	}
	if len(u.FirstName) > 120 || len(u.LastName) > 120 {
		return ErrNameLong
	}
	return nil
}

func (u User) validateEmail() error { return nil }

func (u User) validateAddress() error {
	if len(u.Address) > 120 {
		return ErrAddressLong
	}
	return nil
}

func (u *User) setPassword(password string) error {
	if len(password) < 8 {
		return ErrPasswordShort
	}
	if len(password) > 120 {
		return ErrPasswordLong
	}
	reSmallLetters := regexp.MustCompile("[a-z]")
	reCapital := regexp.MustCompile("[A-Z]")
	reNumbers := regexp.MustCompile("[0-9]")
	smallLettersFind := reSmallLetters.FindAllString(password, -1)
	capitalFind := reCapital.FindAllString(password, -1)
	numberFind := reNumbers.FindAllString(password, -1)
	if len(smallLettersFind) > 0 && len(capitalFind) > 0 && len(numberFind) > 0 {
		return nil
	}
	return ErrPasswordInvalid
}

// SignIn requires both email and password to be passed in and to be checked that they are correct
// before the system would allow access. At the same time, the user model would have its
// lastloginat field updated to the latest timing
func (u *User) SignIn(email, password string) error { return nil }

// Validate checks all entries are correct before passing itself to a service to be saved
func (u User) Validate() error { return nil }

// ForgetPassword resets the forget password token to a random UUID as well as resets the
// forget password expiry token. The function will return the forgetPasswordToken
func (u *User) ForgetPassword() (string, error) { return "", nil }

// ChangePasswordFromForget requires you to provide the forget password token. This function
// will then check the forgetPasswordToken if its correct and alters it accordingly
func (u *User) ChangePasswordFromForget(forgetPasswordToken, password string) error { return nil }

// ChangePassword changes the password on the user object before saving it
func (u *User) ChangePassword(password string) error { return nil }

// ReactivateToken resets the activation token in the case the user did not activate the account
// in time. Returns an activationToken
func (u *User) ReactivateToken() (string, error) { return "", nil }

// Activate user. You would need to provide a activation token to check if it correct.
// If correct, it would return the status of the user which should be true or false
func (u *User) Activate(activationToken string) (bool, error) { return false, nil }

func (a AdminUser) validateStatus() error { return nil }

// Validate function for admin users
func (a AdminUser) Validate() error { return nil }
