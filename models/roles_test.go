package models_test

import (
	"testing"
	"time"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

func TestRolesValidate(t *testing.T) {
	type testCase struct {
		TestName      string
		InputRole     models.Role
		ExpectedError error
	}

	timeCreationTest := time.Now()
	cases := []testCase{}

	for _, singleCase := range cases {
		role := singleCase.InputRole
		err := role.Validate()
		if err == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		}
		if singleCase.ExpectedError == nil && err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		}
		if err != singleCase.ExpectedError {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}

		if role.CreatedAt.After(timeCreationTest) {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "Time needs to be initialized and after initial record time", role.CreatedAt)
		}
		if role.UpdatedAt == time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC) {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "Time needs to be initialized to the 2000-01-01", role.UpdatedAt)
		}
	}
}
