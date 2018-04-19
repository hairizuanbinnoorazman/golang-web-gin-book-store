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
		{TestName: "Last Name Empty", InputFirstName: "acacafacca", ExpectedError: ErrNameShort},
		{TestName: "First Name Short", InputFirstName: "a", InputLastName: "janicanicanioanci", ExpectedError: ErrNameShort},
		{TestName: "Last Name Short", InputLastName: "a", InputFirstName: "janicanicanioanci", ExpectedError: ErrNameShort},
		{TestName: "First Name Long", InputLastName: "ajndjknvjksdnjkvnsjkdnvjkndfjknvjknajkvnjkvsdnjknvjknasvjkdnkjvnakjnvjknsakvnjksdvnkvksdnajkvnkjvdknavkdsnvkdsnkvnsjkdnvskjkvsdankvnkjsnakvnkvdnsjkn", InputFirstName: "janicanicanioanci", ExpectedError: ErrNameLong},
		{TestName: "Last Name Long", InputFirstName: "ajndjknvjksdnjkvnsjkdnvjkndfjknvjknajkvnjkvsdnjknvjknasvjkdnkjvnakjnvjknsakvnjksdvnkvksdnajkvnkjvdknavkdsnvkdsnkvnsjkdnvskjkvsdankvnkjsnakvnkvdnsjkn", InputLastName: "janicanicanioanci", ExpectedError: ErrNameLong},
		{TestName: "Last Name Long First Name Short", InputFirstName: "a", InputLastName: "ajndjknvjksdnjkvnsjkdnvjkndfjknvjknajkvnjkvsdnjknvjknasvjkdnkjvnakjnvjknsakvnjksdvnkvksdnajkvnkjvdknavkdsnvkdsnkvnsjkdnvskjkvsdankvnkjsnakvnkvdnsjkn", ExpectedError: ErrNameShort},
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
		{TestName: "Short Password", Password: "a", ExpectedError: ErrPasswordShort},
		{TestName: "Long Password", Password: "ajsnjfnsdjknvjkdndjknjkdnjkdsnjkdsnjkfndkjfdkjsnjkfdsnjkfdsnknfsdjknfdjknjkafdsnkjdnfknfdanjkfka211io312njk3nkj12njndjknkjdckmkalmklasksjnkjfnskjdnnsdkvj", ExpectedError: ErrPasswordLong},
		{TestName: "Invalid Password Pattern", Password: "aaaaaaaa", ExpectedError: ErrPasswordInvalid},
		{TestName: "Invalid Password Pattern with a number", Password: "aa1aaaaa", ExpectedError: ErrPasswordInvalid},
		{TestName: "Normal Scenario", Password: "Aaaaaaa1", ExpectedError: nil},
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
		{TestName: "Empty Address", Address: "", ExpectedError: nil},
		{TestName: "Long Address", Address: "ajsnjfnsdjknvjkdndjknjkdnjkdsnjkdsnjkfndkjfdkjsnjkfdsnjkfdsnknfsdjknfdjknjkafdsnkjdnfknfdksjnkjfnskjdnnsdkvj", ExpectedError: ErrAddressLong},
		{TestName: "Empty Address", Address: "jkanskdkncajknjkcna cacsc cacaca cs", ExpectedError: nil},
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
		{TestName: "Invalid Pattern Email 1", Email: "aa@aa", ExpectedError: ErrEmailInvalid},
		{TestName: "Invalid Pattern Email 2", Email: "@aa.ca", ExpectedError: ErrEmailInvalid},
		{TestName: "Normal Scenario", Email: "aaaa@aa.ca", ExpectedError: ErrEmailInvalid},
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
