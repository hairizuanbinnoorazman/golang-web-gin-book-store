package models_test

import (
	"strconv"
	"testing"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

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
		}
		if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		}
		if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestForgetPassword(t *testing.T) {
	type testCase struct {
		TestName      string
		User          models.User
		ExpectedError error
	}

	cases := []testCase{}

	for _, singleCase := range cases {
		_, err := singleCase.User.ForgetPassword()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		}
		if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		}
		if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
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

	cases := []testCase{}

	for _, singleCase := range cases {
		token, _ := singleCase.User.ForgetPassword()
		err := singleCase.User.ChangePasswordFromForget(token, singleCase.NewPassword)
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		}
		if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		}
		if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestChangePassword(t *testing.T) {
	type testCase struct {
		TestName      string
		User          models.User
		NewPassword   string
		ExpectedError error
	}

	cases := []testCase{}

	for _, singleCase := range cases {
		err := singleCase.User.ChangePassword(singleCase.NewPassword)
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		}
		if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		}
		if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestReactivateToken(t *testing.T) {
	type testCase struct {
		TestName      string
		User          models.User
		ExpectedError error
	}

	cases := []testCase{}

	for _, singleCase := range cases {
		_, err := singleCase.User.ReactivateToken()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		}
		if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		}
		if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
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

	cases := []testCase{}

	for _, singleCase := range cases {
		status, err := singleCase.User.Activate(singleCase.ActivationToken)
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		}
		if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		}
		if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
		if status != singleCase.ExpectedOutput {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, strconv.FormatBool(singleCase.ExpectedOutput), strconv.FormatBool(singleCase.User.Activated))
		}
	}
}
