package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/seishino/go-example/core/domain"
)

type JWTAuthService struct {
}

// NewJWTAuthService creates a new auth service
func NewJWTAuthService() JWTAuthService {
	return JWTAuthService{}
}

// Authorize authorizes the generated token
func (s JWTAuthService) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("123"), nil
	})
	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired")
		}
	}
	return false, errors.New("couldn't handle token")
}

// CreateToken creates jwt auth token
func (s JWTAuthService) CreateToken(user domain.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Username,
	})

	tokenString, err := token.SignedString([]byte("123"))

	if err != nil {

	}

	return tokenString
}
