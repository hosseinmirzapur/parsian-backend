package main

import (
	"fmt"

	// "github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/common"

	"github.com/joho/godotenv"
)

func main() {
	// app := fiber.New()

	godotenv.Load()

	fmt.Println(common.OrderStatus("jdkjsahdbkjasdbjasdbjkas"))

	// app.Listen(":3000")
}

// func handleError(c *fiber.Ctx, err any) error {
// 	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
// 		"success": false,
// 		"message": err,
// 	})
// }
// func handleSuccess(c *fiber.Ctx, data interface{}) error {
// 	return c.Status(fiber.StatusOK).JSON(data)
// }
