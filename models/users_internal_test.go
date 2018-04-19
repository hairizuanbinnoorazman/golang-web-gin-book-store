package models

import (
	"testing"
)

func TestValidateUserName(t *testing.T) {
	type testCase struct {
		TestName       string
		InputFirstName string
		InputLastName  string
		ExpectedError  error
	}
	cases := []testCase{
		{TestName: "Empty String", InputFirstName: "", InputLastName: "", ExpectedError: ErrNameShort},
		{TestName: "First Name Empty", InputLastName: "acacacac", ExpectedError: ErrNameShort},
	}

	for _, singleCase := range cases {
		tempUser := User{FirstName: singleCase.InputFirstName, LastName: singleCase.InputLastName}
		err := tempUser.validateName()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestSetUserPassword(t *testing.T) {
	type testCase struct {
		TestName      string
		Password      string
		ExpectedError error
	}
	cases := []testCase{
		{TestName: "Empty Password", Password: "", ExpectedError: ErrPasswordShort},
	}

	for _, singleCase := range cases {
		tempUser := User{}
		err := tempUser.setPassword(singleCase.Password)
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestValidateUserAddress(t *testing.T) {
	type testCase struct {
		TestName      string
		Address       string
		ExpectedError error
	}
	cases := []testCase{
		{TestName: "Empty Address", Address: "", ExpectedError: ErrAddressShort},
	}

	for _, singleCase := range cases {
		tempUser := User{Password: singleCase.Address}
		err := tempUser.validateAddress()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestValidateUserEmail(t *testing.T) {
	type testCase struct {
		TestName      string
		Email         string
		ExpectedError error
	}
	cases := []testCase{
		{TestName: "Empty Email", Email: "", ExpectedError: ErrEmailInvalid},
	}

	for _, singleCase := range cases {
		tempUser := User{Password: singleCase.Email}
		err := tempUser.validateEmail()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}
