package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/handlers"
)

func SetupAuth(router fiber.Router) {
	h := handlers.NewAuthHandler()

	router.Post("/login", h.Login)
}
