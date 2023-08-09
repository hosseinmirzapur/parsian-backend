package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/api/helper"
	"github.com/hosseinmirzapur/parsian-backend/services"
)

type orderItemHandler struct{}

func NewOrderItemHandler() *orderItemHandler {
	return &orderItemHandler{}
}

func (h *orderItemHandler) Create(c *fiber.Ctx) error {

	req := new(dto.CreateOrderItemRequest)
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "invalid request",
		})
	}

	id := c.Params("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid url params",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid file upload",
		})
	}
	filepath, err := helper.UploadCtxFile(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "unable to save file",
		})
	}
	req.FilePath = filepath

	_, err = services.CreateOrderItem(req, orderId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "unable to create order-item",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "order-item created successfully",
	})
}
func (h *orderItemHandler) Update(c *fiber.Ctx) error {
	req := new(dto.UpdateOrderItemRequest)
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "invalid request body",
		})
	}
	id := c.Params("id")
	orderItemId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid url params",
		})
	}

	orderItem, err := services.UpdateOrderItem(req, orderItemId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "unable to update order-item",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":    true,
		"order-item": orderItem,
	})
}
func (h *orderItemHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	orderItemId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid url params",
		})
	}
	err = services.DeleteOrderItem(orderItemId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "unable to delete order-item",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Operation Successful",
	})
}
