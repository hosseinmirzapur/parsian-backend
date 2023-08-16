package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() func(c *fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowOrigins: "https://parsian-admin.vercel.app",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Requested-With, Accept-Language, Accept-Encoding, Referer, User-Agent, Cookie",
	})
}
