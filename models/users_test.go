package models_test

import (
	"testing"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

func TestValidate(t *testing.T) {
	type testCase struct {
		TestName       string
		FirstName      string
		LastName       string
		Password       string
		Email          string
		Address        string
		ExpectedOutput string
	}

	cases := []testCase{
		{TestName: "Zero", ExpectedOutput: models.NameShortErr.Error()},
	}

	for _, singleCase := range cases {
		u := models.User{FirstName: singleCase.FirstName,
			LastName: singleCase.LastName,
			Password: singleCase.Password,
			Email:    singleCase.Email,
			Address:  singleCase.Address,
		}
		err := u.Validate()
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
