package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/handlers"
)

func SetupOI(router fiber.Router) {
	h := handlers.NewOrderItemHandler()

	router.Post("/create/:id", h.Create)
	router.Put("/update/:id", h.Update)
	router.Delete("/delete/:id", h.Delete)
}
