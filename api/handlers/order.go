package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/services"
)

type orderHandler struct{}

func NewOrderHandler() *orderHandler {
	return &orderHandler{}
}

func (h *orderHandler) GetAll(c *fiber.Ctx) error {
	orders, err := services.AllOrders()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong getting the orders",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"orders":  orders,
	})
}

func (h *orderHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	orderId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "bad request param",
		})
	}

	order, err := services.GetOrderById(orderId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "something went wrong getting the order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"order":   order,
	})
}
func (h *orderHandler) Create(c *fiber.Ctx) error {
	req := new(dto.CreateOrderRequest)
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": "invalid request body",
		})
	}

	order, err := services.CreateOrder(req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": "unable to create ordre",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"order":   order,
	})
}
func (h *orderHandler) Update(c *fiber.Ctx) error {
	req := new(dto.UpdateOrderRequest)
	id := c.Params("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "bad request param",
		})
	}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": "invalid request body",
		})
	}
	order, err := services.UpdateOrder(req, orderId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "unable to update order",
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"order":   order,
	})
}
func (h *orderHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "invalid request param",
		})
	}

	err = services.DeleteOrder(orderId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "unable to delete order",
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Operation Successful",
	})
}

func (h *orderHandler) GetBySpecialId(c *fiber.Ctx) error {
	specialId := c.Params("specialId")

	order, err := services.FindOrderBySpecialId(specialId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"order":   order,
	})
}

func (h *orderHandler) GetExcelFile(c *fiber.Ctx) error {
	orders, err := services.GetExcelFile()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"excel":   orders,
	})
}
