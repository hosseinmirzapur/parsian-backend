package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/helper"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	godotenv.Load()

	app.Post("/test", func(c *fiber.Ctx) error {
		result, err := helper.UploadToAWS(c)
		if err != nil {
			return handleError(c, err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": result,
		})
	})

	app.Listen(":3000")
}

func handleError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}
