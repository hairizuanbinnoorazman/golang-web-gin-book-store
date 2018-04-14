package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/handlers"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/services"
)

type RouteServices struct {
	Service services.Storage
	User    services.User
}

func setupRouter(rs RouteServices) *gin.Engine {
	router := gin.Default()

	// Auth routes
	router.PUT("/signin")
	router.GET("/activate")
	router.GET("/reactivate")
	router.GET("/forgetpassword", handlers.UserForgetPassword(rs.User))
	router.GET("/confirmforget")

	// User routes
	router.POST("/user", handlers.UserCreate(rs.User))
	router.PUT("/user", handlers.UserUpdate(rs.User))
	router.GET("/user/:id", handlers.UserGet(rs.User))

	return router
}

func main() {
	rs := RouteServices{}
	router := setupRouter(rs)
	router.Run(":8000")
}
