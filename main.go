package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/handlers"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/services"
)

// RouteServices is a struct of interfaces that collates all required services needed by
// the application. This struct is then passed to the setup router function which can then be used
// to alter the actual implementation of the services being run.
//
// Having this allows us to mock the services and test it without hitting databases/external 3rd party APIs
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
	router.GET("/activate", handlers.UserActivate(rs.User))
	router.GET("/reactivate")
	router.GET("/forgetpassword", handlers.UserForgetPassword(rs.User))
	router.POST("/confirmforget", handlers.UserConfirmForget(rs.User))

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
