package models_test

import (
	"testing"
	"time"

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

	timingTest := time.Now()

	for _, singleCase := range cases {
		store, err := models.NewStore()
		if singleCase.ExpectedError == nil && err != nil {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, "nil", err.Error())
		} else if singleCase.ExpectedError != nil && err == nil {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), "nil")
		} else if singleCase.ExpectedError != err {
			t.Errorf("Test Name: %s Expected Output: %s Received output: %s", singleCase.TestName, singleCase.ExpectedError.Error(), err.Error())
		}

		// Check for UUID being set no matter what
		if store.ID != "" {
			t.Errorf("Test Name: UUID check Expected Output: UUID value expected Received output: %s", store.ID)
		}

		if store.CreatedAt.After(timingTest) {
			t.Errorf("Test Name: Created At Expected Output: Created At more than timing test Received output: %s", store.CreatedAt)
		}

		// Check for UUID being set no matter what
		if store.UpdatedAt.After(timingTest) {
			t.Errorf("Test Name: Updated At Expected Output: Created At more than timing tes Received output: %s", store.UpdatedAt)
		}
	}
}
