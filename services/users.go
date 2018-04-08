package services

import "github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"

type User struct {
	Service
}

func (u User) GetByID(ID string) models.User {
	var user models.User
	u.DB.Where("name = ?", "jinzhu").Find(&user)
	return user
}

func (u User) Create(a *models.User) {
	u.DB.Create(a)
}
