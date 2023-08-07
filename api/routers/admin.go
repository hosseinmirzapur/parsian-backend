package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/handlers"
)

func SetupAdmin(router fiber.Router) {
	h := handlers.NewAdminHandler()

	router.Get("/all", h.GetAllAdmins)
	router.Post("/create", h.CreateAdmin)
	router.Put("/update/:id", h.UpdateAdmin)
}
