package users

import (
	"github.com/google/uuid"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

type MockService struct{}

func (u MockService) GetByID(ID string) (models.User, error) {
	a := models.User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockService) GetByEmail(Email string) (models.User, error) {
	a := models.User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockService) Create(*models.User) (models.User, error) {
	a := models.User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockService) Update(*models.User) (models.User, error) {
	a := models.User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}
