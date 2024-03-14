package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userData map[string]string) (string, error) {
	secretKey := []byte(os.Getenv("ACCESS_TOKEN_SECRETKEY"))

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["data"] = userData
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(userId string) (string, error) {
	secretKey := []byte(os.Getenv("REFRESH_TOKEN_SECRETKEY"))

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["data"] = userId
	claims["exp"] = time.Now().AddDate(0, 0, 30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
