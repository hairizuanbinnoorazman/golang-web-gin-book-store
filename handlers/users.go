package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
)

type UserService interface {
	GetByID(ID string) (models.User, error)
	GetByEmail(Email string) (models.User, error)
	Create(*models.User) (models.User, error)
	Update(*models.User) (models.User, error)
}

func UserSignIn(service UserService) func(c *gin.Context) {
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
		c.JSON(http.StatusOK, userUpdated)
	}
}

func UserActivate(service UserService) func(c *gin.Context) {
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

func UserConfirmForget(service UserService) func(c *gin.Context) {
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

func UserCreate(service UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		type Register struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}
		var register Register
		c.BindJSON(&register)
		user := models.NewUser(register.FirstName, register.LastName, register.Email, register.Password)
		_, err := service.Create(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		c.JSON(http.StatusOK, user)
	}
}

func UserGet(service UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := service.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		c.JSON(http.StatusOK, user)
	}
}

func UserUpdate(service UserService) func(c *gin.Context) {
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

func UserForgetPassword(service UserService) func(c *gin.Context) {
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
