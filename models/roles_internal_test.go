package models

import "testing"

func TestRoleValidateName(t *testing.T) {
	type testCase struct {
		TestName      string
		InputRole     Role
		ExpectedError error
	}

	cases := []testCase{}

	for _, singleCase := range cases {
		err := singleCase.InputRole.validateName()
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

func TestRoleValidateStatus(t *testing.T) {
	type testCase struct {
		TestName      string
		InputRole     Role
		ExpectedError error
	}

	cases := []testCase{}

	for _, singleCase := range cases {
		err := singleCase.InputRole.validateStatus()
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
