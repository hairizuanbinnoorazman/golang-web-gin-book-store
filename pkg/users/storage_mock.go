package users

import (
	"github.com/google/uuid"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

type MockStorage struct{}

func (u MockStorage) GetByID(ID string) (models.User, error) {
	a := models.User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockStorage) GetByEmail(Email string) (models.User, error) {
	a := models.User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockStorage) Create(*models.User) (models.User, error) {
	a := models.User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockStorage) Update(*models.User) (models.User, error) {
	a := models.User{ID: uuid.New().String(), FirstName: "John", LastName: "Kelly"}
	return a, nil
}
