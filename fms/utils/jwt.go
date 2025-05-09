package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("s3cr3t")

func CreateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // 2hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(signedToken string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return claims, err
	}

	claims = token.Claims.(jwt.MapClaims)
	return claims, nil
}
