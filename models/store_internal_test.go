package models

import "testing"

func TestStoreValidateName(t *testing.T) {
	type testCase struct {
		TestName      string
		Name          string
		ExpectedError error
	}

	cases := []testCase{
		{TestName: "Empty Name", Name: "", ExpectedError: ErrInvalidRoleName},
		{TestName: "Short Name", Name: "aa", ExpectedError: ErrInvalidRoleName},
		{TestName: "Long Name", Name: "nainaocnjkaajcnjknanjkancjkancjknacjkncajknacjknkjacnkjcanjkcnajknac", ExpectedError: ErrInvalidRoleName},
	}

	for _, singleCase := range cases {
		err := Store{Name: singleCase.Name}.validateStoreName()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestStoreValidateAddress(t *testing.T) {
	type testCase struct {
		TestName      string
		Address       string
		ExpectedError error
	}

	cases := []testCase{}

	for _, singleCase := range cases {
		err := Store{Address: singleCase.Address}.validateStoreAddress()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}
