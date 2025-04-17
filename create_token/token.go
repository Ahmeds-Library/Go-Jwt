package create_token

import (
	"time"

	"github.com/Ahmeds-Library/Go-Jwt/middleware"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		})

	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
