package services

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// ErrJWTSigning is an error thrown during the JWT Signing process
var ErrJWTSigning = errors.New("Error while signing the JWT Token")

// ErrJWTExtract is an error thrown
var ErrJWTExtract = errors.New("Error while extracting values from JWT Token")

// JWTCustomClaims contains the list of values that is to be contained within the JWT Token
type JWTCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// NewToken creates a new JWT token that is to be used by the application
// id generally refer to a string to be stored - in this case is user id
// exp is the expiry in the number of seconds after its issue
// issuer is to specify the application that is providing this token as an identifier
// secret is to seal everything into the token
func NewToken(id string, exp int, secret, issuer string) (string, error) {
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

// ExtractToken takes a JWT Token that is used by the application and extracts the values out
// with the application secret. If process goes well, it would be able to
// return the values
func ExtractToken(tokenString, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*JWTCustomClaims)
	if ok && token.Valid {
		return claims.ID, nil
	}

	return "", ErrJWTExtract
}
