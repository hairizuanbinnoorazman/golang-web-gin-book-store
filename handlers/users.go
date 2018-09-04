package handlers

import (
	"net/http"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/pkg/users"

	"github.com/gin-gonic/gin"
)

type JWTService interface {
	NewToken(id string) (string, error)
	ValidateToken(tokenString string) (string, error)
}

// UserSignIn is a handler function meant to handle user signin
func UserSignIn(service users.Service, jwt JWTService) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Obtain the request details
		type UserSignIn struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var userSignIn UserSignIn
		c.BindJSON(&userSignIn)
		user, errGet := service.GetByEmail(userSignIn.Email)
		if errGet != nil {
			c.JSON(http.StatusUnauthorized, "")
		}
		errSignIn := user.SignIn(userSignIn.Email, userSignIn.Password)
		if errSignIn != nil {
			c.JSON(http.StatusUnauthorized, "")
		}
		userUpdated, errUpdate := service.Update(&user)
		if errUpdate != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		token, errToken := jwt.NewToken(userUpdated.ID)
		if errToken != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		type response struct {
			Token string `json:"string"`
		}
		c.JSON(http.StatusOK, response{Token: token})
	}
}

// UserActivate is a handler function meant to handle activation of users on the platform
func UserActivate(service users.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Obtain the token and id value from the query params
		q := c.Request.URL.Query()
		token := q["token"][0]
		id := q["id"][0]

		user, errGet := service.GetByID(id)
		if errGet != nil {
			c.JSON(http.StatusInternalServerError, "")
		}
		_, errActivate := user.Activate(token)
		if errActivate != nil {
			c.JSON(http.StatusInternalServerError, "")
		}
		c.JSON(http.StatusOK, "")
	}
}

// UserConfirmForget is a handler function meant to handle change of password
// after the token and new password is passed in
func UserConfirmForget(service users.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		type ConfirmForget struct {
			ForgetToken string
			ID          string
			NewPassword string
		}
		var confirmForget ConfirmForget
		c.BindJSON(&confirmForget)
		user, errGet := service.GetByID(confirmForget.ID)
		if errGet != nil {
			c.JSON(http.StatusInternalServerError, "")
		}
		errChange := user.ChangePasswordFromForget(confirmForget.ForgetToken, confirmForget.NewPassword)
		if errChange != nil {
			c.JSON(http.StatusInternalServerError, "")
		}
		c.JSON(http.StatusOK, "")
	}
}

// UserCreate is a handler function meant to handle creation of new users
func UserCreate(service users.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		type Register struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}
		var register Register
		c.BindJSON(&register)
		user, err := users.NewUser(register.FirstName, register.LastName, register.Email, register.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		_, err = service.Create(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, "ancnak")
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// UserGet is a handler function meant to handle creation of getting details of a user
func UserGet(service users.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := service.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		c.JSON(http.StatusOK, user)
	}
}

// UserUpdate is a handler function meant to handle update of the details of a user
func UserUpdate(service users.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		type userUpdate struct {
			ID        string
			FirstName string
			LastName  string
			Address   string
		}
		var updateValues userUpdate
		c.BindJSON(updateValues)
		user, err := service.GetByID(updateValues.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		user.FirstName = updateValues.FirstName
		user.LastName = updateValues.LastName
		user.Address = updateValues.Address
		_, errUpdate := service.Update(&user)
		if errUpdate != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		c.JSON(http.StatusOK, user)
	}
}

// UserForgetPassword is a handler function meant to initiate the creation of
// forget password tokens etc
func UserForgetPassword(service users.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		type userForget struct {
			Email string `json:"email"`
		}
		var forget userForget
		c.BindJSON(forget)
		user, err := service.GetByEmail(forget.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		_, errToken := user.ForgetPassword()
		if errToken != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		c.JSON(http.StatusOK, "")
	}
}
