package services

import (
	"time"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

type User struct {
	Storage
}

func (u User) GetByID(ID string) (models.User, error) {
	var user models.User
	// u.DB.Where("name = ?", "jinzhu").Find(&user)
	user.FirstName = "LOL"
	user.LastName = "MIAO"
	return user, nil
}

func (u User) Create(a *models.User) (models.User, error) {
	// u.DB.Create(a)
	lol := models.User{FirstName: "kcmkamcklamlc", LastLoginAt: time.Now()}
	return lol, nil
}

func (u User) Update(a *models.User) (models.User, error) {
	lol := models.User{FirstName: "kcmkamcklamlc", LastLoginAt: time.Now()}
	return lol, nil
}

func (u User) GetByEmail(a string) (models.User, error) {
	lol := models.User{FirstName: "kcmkamcklamlc", LastLoginAt: time.Now()}
	return lol, nil
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
