package main

// "fmt"

// "github.com/go-playground/validator/v10"
// "github.com/gofiber/fiber/v2"
// "github.com/hosseinmirzapur/parsian-backend/api/dto"
// "github.com/hosseinmirzapur/parsian-backend/api/helper"
// "github.com/joho/godotenv"

// func main() {
// 	app := fiber.New()

// 	godotenv.Load()

// 	app.Post("/test", func(c *fiber.Ctx) error {
// 		result, err := helper.UploadToAWS(c)
// 		if err != nil {
// 			return handleError(c, err)
// 		}

// 		return c.Status(200).JSON(&fiber.Map{
// 			"success": true,
// 			"message": result,
// 		})
// 	})

// 	app.Post("/validation", func(c *fiber.Ctx) error {
// 		var err error
// 		req := new(dto.CreateOrderRequest)
// 		err = c.BodyParser(req)
// 		if err != nil {
// 			fmt.Println("body parser error")
// 			return handleError(c, err)
// 		}
// 		err = req.Validate(req)

// 		if err != nil {
// 			validationErrors := []map[string]string{}

// 			for _, err := range err.(validator.ValidationErrors) {
// 				validationErrors = append(validationErrors, map[string]string{
// 					err.Field(): err.Tag(),
// 				})
// 			}

// 			return handleError(c, validationErrors)
// 		}

// 		return handleSuccess(c, &fiber.Map{
// 			"message": "worked",
// 			"request": req,
// 		})

// 	})

// 	app.Listen(":3000")
// }

// func handleError(c *fiber.Ctx, err any) error {
// 	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
// 		"success": false,
// 		"message": err,
// 	})
// }
// func handleSuccess(c *fiber.Ctx, data interface{}) error {
// 	return c.Status(fiber.StatusOK).JSON(data)
// }
