package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/services"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	// Validation Section
	var err error
	req := new(dto.LoginRequest)
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	// Business Logic Section
	token, admin, err := services.AdminLogin(req.Username, req.Password)

	if err != nil && token != "" {
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
