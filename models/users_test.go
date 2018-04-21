package models_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

func TestNewUser(t *testing.T) {
	type testCase struct {
		TestName      string
		FirstName     string
		LastName      string
		Password      string
		Email         string
		ExpectedError error
	}

	cases := []testCase{
		{TestName: "Empty Values", ExpectedError: models.ErrPasswordShort},
		{TestName: "No Password", FirstName: "aaaa", LastName: "bbbb", Email: "aaaa@aa.ca", ExpectedError: models.ErrPasswordShort},
		{TestName: "Normal Scenario", FirstName: "aaaa", LastName: "bbbb", Email: "aaaaaa@ca.ea", Password: "maskdkd1231a", ExpectedError: models.ErrPasswordInvalid},
		{TestName: "Normal Scenario", FirstName: "aaaa", LastName: "bbbb", Email: "aaaaaa@ca.ea", Password: "maskdkd1231A", ExpectedError: nil},
	}

	for _, singleCase := range cases {
		u, errPass := models.NewUser(singleCase.FirstName, singleCase.LastName, singleCase.Email, singleCase.Password)
		if errPass != nil {
			if singleCase.ExpectedError == nil {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", errPass.Error())
			} else if singleCase.ExpectedError != errPass {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), errPass.Error())
			}
		}
		if u != nil {
			err := u.Validate()
			if err == nil && err != singleCase.ExpectedError {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
			} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
			} else if err != singleCase.ExpectedError {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
			}
			if u.ID == "" {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "A UUID Value to be provided for user id", u.ID)
			}
			if u.ActivationToken == "" {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "A UUID Value to be provided for activation token", u.ID)
			}
			if u.Password == singleCase.Password {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "Password values needs to be encrypted", u.Password)
			}
		}
	}
}

func TestSignIn(t *testing.T) {
	type testCase struct {
		TestName      string
		LoginEmail    string
		LoginPassword string
		ExpectedError error
	}

	// Set a account that is validated correctly
	u, _ := models.NewUser("aaaa", "aaavv", "aaaaa@sac.ca", "acaockoasc1mcaA")
	u.SignIn("aaaaa@sac.ca", "acaockoasc1mcaA")
	initialLoginTime := u.LastLoginAt.Add(-5 * time.Second)

	cases := []testCase{
		{TestName: "Empty Value", ExpectedError: models.ErrLogin},
		{TestName: "Empty Password", LoginEmail: "aacac@caca.ac", ExpectedError: models.ErrLogin},
		{TestName: "Wrong Password", LoginEmail: "aaaaa@sac.ca", LoginPassword: "acaockoasc2mcaA", ExpectedError: models.ErrLogin},
		{TestName: "Normal Scenario", LoginEmail: "aaaaa@sac.ca", LoginPassword: "acaockoasc1mcaA", ExpectedError: nil},
	}

	for _, singleCase := range cases {
		err := u.SignIn(singleCase.LoginEmail, singleCase.LoginPassword)
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
		// Check to make sure that last login time is updated
		if err == singleCase.ExpectedError {
			if !u.LastLoginAt.After(initialLoginTime) {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, initialLoginTime, u.LastLoginAt)
			}
		}
	}
}

func TestUserValidate(t *testing.T) {
	type testCase struct {
		TestName      string
		FirstName     string
		LastName      string
		Password      string
		Email         string
		Address       string
		ExpectedError error
	}

	cases := []testCase{
		{TestName: "Zero", ExpectedError: models.ErrNameShort},
		{TestName: "Missing Last Name", FirstName: "Lol", ExpectedError: models.ErrNameShort},
		{TestName: "Missing First Name", LastName: "Miao", ExpectedError: models.ErrNameShort},
		{TestName: "Email and password missing", FirstName: "Laaol", LastName: "Miao", ExpectedError: models.ErrEmailInvalid},
		{TestName: "Email missing", FirstName: "Lolaa", LastName: "Miaaao", Password: "acac", ExpectedError: models.ErrEmailInvalid},
		{TestName: "All Filled Up accordingly", FirstName: "Loal", LastName: "Miao", Password: "acaaaaac", Email: "aaa@aaa.aa", ExpectedError: nil},
	}

	for _, singleCase := range cases {
		u := models.User{FirstName: singleCase.FirstName,
			LastName: singleCase.LastName,
			Password: singleCase.Password,
			Email:    singleCase.Email,
			Address:  singleCase.Address,
		}
		err := u.Validate()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestForgetPassword(t *testing.T) {
	type testCase struct {
		TestName string
		User     models.User
	}

	cases := []testCase{
		{TestName: "Normal Scenario", User: models.User{}},
	}

	for _, singleCase := range cases {
		token, _ := singleCase.User.ForgetPassword()
		if token == "" {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "UUID Value", token)
		}
	}
}

func TestChangePasswordFromForget(t *testing.T) {
	type testCase struct {
		TestName      string
		User          models.User
		NewPassword   string
		ExpectedError error
	}

	cases := []testCase{
		{TestName: "Normal Scenario", User: models.User{}, NewPassword: "aa1Aaaaaa", ExpectedError: nil},
		{TestName: "Short Password", User: models.User{}, NewPassword: "aaaaaa", ExpectedError: models.ErrPasswordShort},
		{TestName: "Same Password", User: models.User{Password: "aaa1Aaaaa"}, NewPassword: "aaa1Aaaaa", ExpectedError: nil},
	}

	for _, singleCase := range cases {
		token, _ := singleCase.User.ForgetPassword()
		singleCase.User.ForgetPasswordToken = token
		err := singleCase.User.ChangePasswordFromForget(token, singleCase.NewPassword)
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestChangePassword(t *testing.T) {
	type testCase struct {
		TestName      string
		OldPassword   string
		NewPassword   string
		ExpectedError error
	}

	cases := []testCase{
		{TestName: "Normal Scenario", OldPassword: "a1Aaaaaaa", NewPassword: "bbb1Bbbbb", ExpectedError: nil},
		{TestName: "Same Password", OldPassword: "1aA2345678", NewPassword: "1aA2345678", ExpectedError: models.ErrSamePassword},
	}

	for _, singleCase := range cases {
		user, _ := models.NewUser("aaaa", "aaaa", "aaaa@aa.aa", singleCase.OldPassword)
		err := user.ChangePassword(singleCase.NewPassword)
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestReactivateToken(t *testing.T) {
	type testCase struct {
		TestName string
		User     models.User
	}

	cases := []testCase{
		{TestName: "Common Scenario", User: models.User{}},
	}

	for _, singleCase := range cases {
		initialToken := singleCase.User.ActivationToken
		token, _ := singleCase.User.ReactivateToken()
		if token == "" {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "UUID Value Expected", token)
		}
		if singleCase.User.ActivationToken == "" {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "UUID Value Expected", token)
		}
		if initialToken == token {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "Different Values Expected", token)
		}
	}
}

func TestActivate(t *testing.T) {
	type testCase struct {
		TestName        string
		User            models.User
		ActivationToken string
		ExpectedOutput  bool
		ExpectedError   error
	}

	cases := []testCase{
		{TestName: "Normal Scenario", User: models.User{ActivationToken: "abc"}, ActivationToken: "abc", ExpectedOutput: true, ExpectedError: nil},
		{TestName: "Wrong Token", User: models.User{ActivationToken: "abc123"}, ActivationToken: "abc", ExpectedOutput: false, ExpectedError: models.ErrActivationTokenInvalid},
	}

	for _, singleCase := range cases {
		status, err := singleCase.User.Activate(singleCase.ActivationToken)
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
		if status != singleCase.ExpectedOutput {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, strconv.FormatBool(singleCase.ExpectedOutput), strconv.FormatBool(singleCase.User.Activated))
		}
	}
}
