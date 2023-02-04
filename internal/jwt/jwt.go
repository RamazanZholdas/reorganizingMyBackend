package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func NewJwtTokenWithClaims(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    email,
		ExpiresAt: &jwt.Time{Time: time.Now().Add(time.Hour * 24)},
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
