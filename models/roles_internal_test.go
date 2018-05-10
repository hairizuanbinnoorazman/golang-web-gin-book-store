package models

import "testing"

func TestRoleValidateName(t *testing.T) {
	type testCase struct {
		TestName      string
		InputRole     Role
		ExpectedError error
	}

	cases := []testCase{
		{TestName: "Empty Name", InputRole: Role{Name: ""}, ExpectedError: ErrInvalidRoleName},
		{TestName: "Short Name", InputRole: Role{Name: "aa"}, ExpectedError: ErrInvalidRoleName},
		{TestName: "Admin case", InputRole: Role{Name: "storeadmin"}, ExpectedError: nil},
		{TestName: "Managerial case", InputRole: Role{Name: "manager"}, ExpectedError: nil},
		{TestName: "Supplier case", InputRole: Role{Name: "supplier"}, ExpectedError: nil},
		{TestName: "Supplier case", InputRole: Role{Name: "Admin"}, ExpectedError: ErrInvalidRoleName},
		{TestName: "Supplier case", InputRole: Role{Name: "StoreAdmin"}, ExpectedError: ErrInvalidRoleName},
	}

	for _, singleCase := range cases {
		err := singleCase.InputRole.validateName()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
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

	cases := []testCase{
		{TestName: "Empty Status", InputRole: Role{Status: ""}, ExpectedError: ErrInvalidStatus},
		{TestName: "Active Status", InputRole: Role{Status: "active"}, ExpectedError: nil},
		{TestName: "Deactivated Status", InputRole: Role{Status: "deactivated"}, ExpectedError: nil},
		{TestName: "Depreciated Status", InputRole: Role{Status: "depreciated"}, ExpectedError: nil},
		{TestName: "Active Status", InputRole: Role{Status: "Active"}, ExpectedError: ErrInvalidStatus},
	}

	for _, singleCase := range cases {
		err := singleCase.InputRole.validateStatus()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}

func TestRoleValidateDescription(t *testing.T) {
	type testCase struct {
		TestName      string
		InputRole     Role
		ExpectedError error
	}

	cases := []testCase{}

	for _, singleCase := range cases {
		err := singleCase.InputRole.validateDescription()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}
	}
}
