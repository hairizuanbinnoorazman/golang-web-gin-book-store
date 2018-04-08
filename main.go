package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/handlers"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/services"
)

type RouteServices struct {
	User services.User
}

func setupRouter(rs RouteServices) *gin.Engine {
	router := gin.Default()
	router.GET("/welcome", welcome)
	router.GET("/user/:id", handlers.UserGet(rs.User))
	return router
}

func main() {
	router := setupRouter()
	router.Run(":8000")
}

func welcome(c *gin.Context) {
	type Lol struct {
		Status string
	}
	c.JSON(http.StatusOK, Lol{Status: "ok"})
}
