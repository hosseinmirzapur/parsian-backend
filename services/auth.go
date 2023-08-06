package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(adminId string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	secret := os.Getenv("JWT_SECRET")
	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = adminId
	claims["exp"] = now.Add(time.Hour * 24).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
