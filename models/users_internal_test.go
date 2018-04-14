package models

import (
	"testing"
)

func TestValidateName(t *testing.T) {
	type testCase struct {
		TestName       string
		InputFirstName string
		InputLastName  string
		ExpectedOutput string
	}
	cases := []testCase{
		{TestName: "Empty String", InputFirstName: "", InputLastName: "", ExpectedOutput: "Name cannot be shorter than 8 characters"},
	}

	for _, singleCase := range cases {
		tempUser := User{FirstName: singleCase.InputFirstName, LastName: singleCase.InputLastName}
		err := tempUser.validateName()
		if err != nil {
			if err.Error() != singleCase.ExpectedOutput {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedOutput, err.Error())
			}
		} else {
			if singleCase.ExpectedOutput != "nil" {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedOutput, "nil")
			}
		}
	}
}

func TestSetPassword(t *testing.T) {
	type testCase struct {
		TestName       string
		Password       string
		ExpectedOutput string
	}
	cases := []testCase{
		{TestName: "Empty Password", Password: "", ExpectedOutput: "Password is too short"},
	}

	for _, singleCase := range cases {
		tempUser := User{}
		err := tempUser.setPassword(singleCase.Password)
		if err != nil {
			if err.Error() != singleCase.ExpectedOutput {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedOutput, err.Error())
			}
		} else {
			if singleCase.ExpectedOutput != "nil" {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedOutput, "nil")
			}
		}
	}
}

func TestValidateAddress(t *testing.T) {
	type testCase struct {
		TestName       string
		Address        string
		ExpectedOutput string
	}
	cases := []testCase{
		{TestName: "Empty Password", Address: "", ExpectedOutput: "Address is too short"},
	}

	for _, singleCase := range cases {
		tempUser := User{Password: singleCase.Address}
		err := tempUser.validateAddress()
		if err != nil {
			if err.Error() != singleCase.ExpectedOutput {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedOutput, err.Error())
			}
		} else {
			if singleCase.ExpectedOutput != "nil" {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedOutput, "nil")
			}
		}
	}
}

func TestValidateEmail(t *testing.T) {
	type testCase struct {
		TestName       string
		Email          string
		ExpectedOutput string
	}
	cases := []testCase{
		{TestName: "Empty Password", Email: "", ExpectedOutput: "Email is invalid"},
	}

	for _, singleCase := range cases {
		tempUser := User{Password: singleCase.Email}
		err := tempUser.validateEmail()
		if err != nil {
			if err.Error() != singleCase.ExpectedOutput {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedOutput, err.Error())
			}
		} else {
			if singleCase.ExpectedOutput != "nil" {
				t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedOutput, "nil")
			}
		}
	}
}
