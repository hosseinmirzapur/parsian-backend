package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/data/db"
	"github.com/hosseinmirzapur/parsian-backend/data/models"
	"github.com/hosseinmirzapur/parsian-backend/services"
)

type adminHandler struct{}

func NewAdminHandler() *adminHandler {
	return &adminHandler{}
}

func (h *adminHandler) GetAllAdmins(c *fiber.Ctx) error {
	var admins []models.Admin
	var err error

	dbClient := db.GetDB()
	err = dbClient.Find(&admins).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to get admins",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"admins":  admins,
	})
}
func (h *adminHandler) CreateAdmin(c *fiber.Ctx) error {
	req := new(dto.CreateAdminRequest)

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	admin, err := services.CreateAdmin(req.Username, req.Password, req.Name, req.Role)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Unable to create admin",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"admin":   admin,
	})
}

func (h *adminHandler) UpdateAdmin(c *fiber.Ctx) error {

	req := new(dto.UpdateAdminRequest)
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	id := c.Params("id")

	adminId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid url param",
		})
	}

	admin, err := services.UpdateAdmin(req, uint(adminId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Unable to update admin",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"admin":   admin,
	})
}

func (h *adminHandler) DeleteAdmin(c *fiber.Ctx) error {
	id := c.Params("id")

	adminId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Bad request param",
		})
	}

	err = services.DeleteAdmin(uint(adminId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "unable to delete admin",
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Operation Successful",
	})
}
