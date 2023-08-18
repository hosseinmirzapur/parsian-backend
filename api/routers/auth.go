package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/handlers"
	"github.com/hosseinmirzapur/parsian-backend/api/middlewares"
)

func SetupAuth(router fiber.Router) {
	h := handlers.NewAuthHandler()

	router.Post("/login", h.Login)
	router.Get("/check", middlewares.ProtectedRoute(), h.IsLoggedIn)
}
