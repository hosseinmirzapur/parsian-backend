// package main

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/hosseinmirzapur/parsian-backend/api/dto"
// 	"github.com/hosseinmirzapur/parsian-backend/utils"
// 	"github.com/xuri/excelize/v2"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	app := fiber.New()

// 	godotenv.Load()

// 	app.Post("/form", func(c *fiber.Ctx) error {
// 		req, err := dto.ValidateCreateOrderForm(c)

// 		if err != nil {
// 			return handleError(c, err)
// 		}

// 		return handleSuccess(c, req)
// 	})

// 	app.Get("/excel", func(c *fiber.Ctx) error {
// 		f := excelize.NewFile()

// 		defer func() {
// 			if err := f.Close(); err != nil {
// 				handleError(c, err)
// 			}
// 		}()

// 		sheet := "Sheet1"

// 		index, err := f.NewSheet(sheet)

// 		if err != nil {
// 			return handleError(c, err)
// 		}

// 		f.SetActiveSheet(index)

// 		f.SetCellValue(sheet, "A1", "Hello World 1")
// 		f.SetCellValue(sheet, "A2", "Hello World 2")
// 		f.SetCellValue(sheet, "A3", "Hello World 3")

// 		if err := f.SaveAs("export.xlsx"); err != nil {
// 			return handleError(c, err)
// 		}

// 		dPath, err := utils.UploadToAWS("export.xlsx")

// 		if err != nil {
// 			return handleError(c, err)
// 		}

// 		return handleSuccess(c, &fiber.Map{
// 			"link": dPath,
// 		})
// 	})

// 	app.Listen(":3000")
// }

//	func handleError(c *fiber.Ctx, err any) error {
//		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
//			"success": false,
//			"message": err.(error).Error(),
//		})
//	}
//
//	func handleSuccess(c *fiber.Ctx, data interface{}) error {
//		return c.Status(fiber.StatusOK).JSON(data)
//	}
package main
