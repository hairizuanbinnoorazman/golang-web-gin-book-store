package models_test

import (
	"testing"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

func TestValidate(t *testing.T) {
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
}

func TestReactivateToken(t *testing.T) {}

func TestActivate(t *testing.T) {}
