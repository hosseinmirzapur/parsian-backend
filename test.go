package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/dto"

	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	godotenv.Load()

	app.Post("/form", func(c *fiber.Ctx) error {
		req, err := dto.ValidateCreateOrderForm(c)

		if err != nil {
			return handleError(c, err)
		}

		return handleSuccess(c, req)
	})

	app.Listen(":3000")
}

func handleError(c *fiber.Ctx, err any) error {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"success": false,
		"message": err.(error).Error(),
	})
}
func handleSuccess(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}
