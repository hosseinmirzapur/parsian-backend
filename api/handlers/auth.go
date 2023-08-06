package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var err error
	req := new(dto.LoginRequest)
	err = c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	return nil // it's gonna change in the future...
}
