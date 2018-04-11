package services

import (
	"time"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

type User struct {
	Storage
}

func (u User) GetByID(ID string) models.User {
	var user models.User
	// u.DB.Where("name = ?", "jinzhu").Find(&user)
	user.FirstName = "LOL"
	user.LastName = "MIAO"
	return user
}

func (u User) Create(a *models.User) models.User {
	// u.DB.Create(a)
	lol := models.User{FirstName: "kcmkamcklamlc", LastLoginAt: time.Now()}
	return lol
}

type User2 struct {
	TestData models.User
}

func (u User2) GetByID(ID string) models.User {
	u.TestData.FirstName = "lol"
	u.TestData.LastName = "cacacarjnejkwkf"
	return u.TestData
}

func (u User2) Create(a *models.User) models.User {
	return *a
}
