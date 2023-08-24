package dto

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateOrderItemRequest struct {
	Name             string `json:"name" validation:"required"`
	AllowSandPaper   bool   `json:"allow_sand_paper" validation:"required"`
	AllowDestruction bool   `json:"allow_destruction" validation:"required"`
	Status           string `json:"status" validation:"required,oneof=pending partial office paid"`
	TestType         string `json:"test_type" validation:"required,oneof=analyze hardness both"`
	Quantity         uint   `json:"quantity" validation:"required,min=1"`
	Description      string `json:"description"`
}

type UpdateOrderItemRequest struct {
	Name             string `json:"name,omitempty"`
	AllowSandPaper   bool   `json:"allow_sand_paper,omitempty"`
	AllowDestruction bool   `json:"allow_destruction,omitempty"`
	Status           string `json:"status,omitempty" validation:"oneof=pending partial office paid"`
	TestType         string `json:"test_type,omitempty" validation:"oneof=analyze hardness both"`
	Quantity         uint   `json:"quantity,omitempty"`
	Description      string `json:"description,omitempty"`
}

func ValidateCreateOrderForm(c *fiber.Ctx) (CreateOrderItemRequest, error) {
	hasError := false
	name := c.FormValue("name")
	allowSandPaper := c.FormValue("allow_sand_paper")
	allowDestruction := c.FormValue("allow_destruction")
	status := c.FormValue("status")
	testType := c.FormValue("test_type")
	quantity := c.FormValue("quantity")
	hasError = isEmpty(name) || isEmpty(allowSandPaper) || isEmpty(allowDestruction) || isEmpty(status) || isEmpty(testType) || isEmpty(quantity)

	// Required fields
	if hasError {
		return CreateOrderItemRequest{}, fmt.Errorf("invalid form data")
	}

	// Validate Order Status
	err := validateStatus(status)
	if err != nil {
		return CreateOrderItemRequest{}, err
	}

	// Validate Test Types
	err = validateTestType(testType)
	if err != nil {
		return CreateOrderItemRequest{}, err
	}

	quantityInt, err := strconv.Atoi(quantity)

	if err != nil {
		return CreateOrderItemRequest{}, err
	}
	req := CreateOrderItemRequest{
		Name:             name,
		AllowSandPaper:   allowSandPaper == "1" || allowSandPaper == "true",
		AllowDestruction: allowDestruction == "1" || allowDestruction == "true",
		Status:           status,
		TestType:         testType,
		Quantity:         uint(quantityInt),
		Description:      "",
	}

	return req, nil
}

func isEmpty(value string) bool {
	return value == ""
}

func validateStatus(status string) error {
	if status != "pending" && status != "partial" && status != "office" && status != "paid" {
		return fmt.Errorf("invalid status")
	}
	return nil
}

func validateTestType(testType string) error {
	if testType != "analyze" && testType != "hardness" && testType != "both" {
		return fmt.Errorf("invalid test type")
	}
	return nil
}
