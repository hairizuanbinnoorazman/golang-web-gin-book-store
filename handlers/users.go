package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

type UserService interface {
	GetByID(ID string) models.User
	Create(*models.User) models.User
}

func UserCreate(service UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		type SignIn struct {
			FirstName string
			LastName  string
			Email     string
			Password  string
		}
		var signin SignIn
		c.BindJSON(&signin)
		user := models.User{FirstName: signin.FirstName, LastName: signin.LastName, Password: signin.Password, Email: signin.Email}
		service.Create(&user)
		c.JSON(http.StatusOK, user)
	}
}

func UserGet(service UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		user := service.GetByID(id)
		c.JSON(http.StatusOK, user)
	}
}
