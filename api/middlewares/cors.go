package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() func(c *fiber.Ctx) error {
	return cors.New(cors.Config{
		// AllowOrigins: "https://parsian-backend.vercel.app",
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	})
}
