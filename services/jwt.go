package services

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var ErrJWTSigning = errors.New("Error while signing the JWT Token")
var ErrJWTValidate = errors.New("Error while validating JWT Token")

type JWTCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// NewToken creates a new JWT token that is to be used by the application
func NewToken(id string) (string, error) {
	exp := viper.GetInt("jwt_expire")
	secret := viper.GetString("jwt_secret")
	issuer := viper.GetString("jwt_issuer")

	claims := JWTCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Duration(exp) * time.Second).Unix()),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", ErrJWTSigning
	}

	return tokenString, nil
}

// ValidateToken takes a JWT Token that is used by the application and validates it
// with the application secret. If validation process goes well, it would be able to
// return the values intended
func ValidateToken(tokenString string) (*JWTCustomClaims, error) {
	// Retrive secret for retriving values
	secret := viper.GetString("jwt_secret")

	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil {
		return &JWTCustomClaims{}, err
	}

	claims, ok := token.Claims.(*JWTCustomClaims)
	if ok && token.Valid {
		return &JWTCustomClaims{
			ID: claims.ID,
		}, nil
	}

	return &JWTCustomClaims{}, ErrJWTValidate
}
