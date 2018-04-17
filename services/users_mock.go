package services

import "github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"

type MockUser struct{}

func (u MockUser) GetByID(ID string) (models.User, error) {
	a := models.User{FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockUser) GetByEmail(Email string) (models.User, error) {
	a := models.User{FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockUser) Create(*models.User) (models.User, error) {
	a := models.User{FirstName: "John", LastName: "Kelly"}
	return a, nil
}

func (u MockUser) Update(*models.User) (models.User, error) {
	a := models.User{FirstName: "John", LastName: "Kelly"}
	return a, nil
}
