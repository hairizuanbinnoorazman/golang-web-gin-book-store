package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/handlers"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/services"
)

type RouteServices struct {
	User handlers.UserService
}

func setupRouter(rs RouteServices) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsHandler := cors.New(
		cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowAllOrigins:  true,
		})
	router.Use(corsHandler)

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
	rs := RouteServices{User: services.User{}}
	router := setupRouter(rs)
	router.Run(":8000")
}
