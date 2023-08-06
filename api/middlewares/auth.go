package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedRoute() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		bearerToken := c.Get("Authorization")[7:]

		if bearerToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token, _ := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if time.Now().Unix() > claims["exp"].(int64) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Token Expired",
				})
			}
			c.Locals("admin_id", claims["admin_id"])

		} else {
			return c.Status(fiber.ErrExpectationFailed.Code).JSON(fiber.Map{
				"message": fiber.ErrExpectationFailed.Message,
			})
		}

		return c.Next()

	}
}
