package services

import (
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

type User struct {
	Storage
}

func (u User) GetByID(ID string) (models.User, error) {
	tempUser := models.User{ID: ID}
	result := u.Storage.DB.First(&tempUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return tempUser, nil
}

func (u User) Create(a *models.User) (models.User, error) {
	result := u.Storage.DB.Create(a)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return *a, nil
}

func (u User) Update(a *models.User) (models.User, error) {
	result := u.Storage.DB.Update(a)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return *a, nil
}

func (u User) GetByEmail(a string) (models.User, error) {
	tempUser := models.User{}
	result := u.Storage.DB.Where("email = ?", a).First(&tempUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return tempUser, nil
}

type User2 struct {
	TestData models.User
}

func (u User2) GetByID(ID string) (models.User, error) {
	u.TestData.FirstName = "lol"
	u.TestData.LastName = "cacacarjnejkwkf"
	return u.TestData, nil
}

func (u User2) Create(a *models.User) (models.User, error) {
	return *a, nil
}
