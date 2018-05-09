package models_test

import (
	"testing"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

func TestNewStore(t *testing.T) {
	type testCase struct {
		TestName      string
		Name          string
		ExpectedError error
	}

	cases := []testCase{
		{TestName: "Empty store", ExpectedError: models.ErrStoreNameInvalid},
	}

	for _, singleCase := range cases {
		_, err := models.NewStore()
		if singleCase.ExpectedError == nil && err != nil {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if singleCase.ExpectedError != nil && err == nil {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError != err {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}

		// Check for UUID being set no matter what
	}
}
