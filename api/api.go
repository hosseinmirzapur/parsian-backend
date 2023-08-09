package api

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/middlewares"
	"github.com/hosseinmirzapur/parsian-backend/api/routers"
)

func InitServer() {
	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	})

	app.Use(middlewares.Cors())
	RegisterRoutes(app)

	app.Listen(":3000")
}

func RegisterRoutes(app *fiber.App) {
	// Serve Static Files
	app.Static("/", "./public")

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
