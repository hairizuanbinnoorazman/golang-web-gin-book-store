package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware1(gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Before", "cacaca")
		c.Next()
		if len(c.Errors) > 0 {
			c.AbortWithError(http.StatusBadRequest, errors.New("Oh shitos"))
		}
		c.Header("X-After", "cacaca")
	}
}
