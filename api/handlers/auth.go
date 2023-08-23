package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	validation "github.com/hosseinmirzapur/parsian-backend/api/validations"
	"github.com/hosseinmirzapur/parsian-backend/services"
)

type authHandler struct {
}

func NewAuthHandler() *authHandler {
	return &authHandler{}
}

func (h *authHandler) Login(c *fiber.Ctx) error {

	// Validation Section
	var err error
	req := new(dto.LoginRequest)
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	errs, ok := validation.ValidateData(req)

	if !ok {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": errs,
		})
	}

	// Business Logic Section
	token, admin, err := services.AdminLogin(req.Username, req.Password)

	if err != nil || token == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Wrong username or password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"token":   token,
		"admin":   admin,
	}) // it's gonna change in the future...
}

func (h *authHandler) IsLoggedIn(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"loggedIn": true,
	})
}
