package api

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hosseinmirzapur/parsian-backend/api/middlewares"
	"github.com/hosseinmirzapur/parsian-backend/api/routers"
)

func InitServer() {
	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	})
	// app.Use(logger.New())

	app.Use(middlewares.Cors())

	RegisterRoutes(app)

	port := os.Getenv("PORT")

	app.Listen(fmt.Sprintf(":%s", port))
}

func RegisterRoutes(app *fiber.App) {

	// Api Related Stuff
	api := app.Group("/api")

	v1 := api.Group("/v1")
	{
		// Auth
		authRouter := v1.Group("/auth")
		routers.SetupAuth(authRouter)

		// Admin
		adminRouter := v1.Group("/admin", middlewares.ProtectedRoute())
		routers.SetupAdmin(adminRouter)

		// Order
		orderRouter := v1.Group("/order", middlewares.ProtectedRoute())
		routers.SetupOrder(orderRouter)

		// Order-Item
		oiRouter := v1.Group("/oi", middlewares.ProtectedRoute())
		routers.SetupOI(oiRouter)

	}

}
