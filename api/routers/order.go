package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/handlers"
)

func SetupOrder(router fiber.Router) {
	h := handlers.NewOrderHandler()

	router.Get("/all", h.GetAll)
	router.Get("/:id", h.GetById)
	router.Post("/create", h.Create)
	router.Put("/update/:id", h.Update)
	router.Delete("/delete/:id", h.Delete)
	router.Get("/special/:specialId", h.GetBySpecialId)
	router.Post("/excel", h.GetExcelFile)
}
