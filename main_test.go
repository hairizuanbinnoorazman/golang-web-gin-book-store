package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/services"
	"github.com/magiconair/properties/assert"
)

func TestUserRoutes(t *testing.T) {
	rs := RouteServices{User: services.MockUser{}}
	router := setupRouter(rs)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/12", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}
