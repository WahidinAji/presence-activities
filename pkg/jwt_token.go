package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func JWTToken(identity string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = identity
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	jwt, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}
	return jwt, nil
}
