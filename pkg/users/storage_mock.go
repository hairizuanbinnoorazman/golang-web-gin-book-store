package users

import (
	"github.com/google/uuid"
)

type MockService struct{}

func (u MockService) GetByID(ID string) (User, error) {
	a := User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockService) GetByEmail(Email string) (User, error) {
	a := User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockService) Create(*User) (User, error) {
	a := User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockService) Update(*User) (User, error) {
	a := User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}
