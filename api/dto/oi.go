package dto

import (
	"github.com/hosseinmirzapur/parsian-backend/common"
)

type CreateOrderItemRequest struct {
	Name             string             `json:"name" validation:"required"`
	AllowSandPaper   bool               `json:"allow_sand_paper" validation:"required"`
	AllowDestruction bool               `json:"allow_destruction" validation:"required"`
	Status           common.OrderStatus `json:"status" validation:"required"`
	TestType         common.TestType    `json:"test_type" validation:"required"`
	Quantity         uint               `json:"quantity" validation:"required"`
	Description      string             `json:"description" validation:"required"`
}

type UpdateOrderItemRequest struct {
	Name             string `json:"name,omitempty"`
	AllowSandPaper   bool   `json:"allow_sand_paper,omitempty"`
	AllowDestruction bool   `json:"allow_destruction,omitempty"`
	Status           string `json:"status,omitempty"`
	TestType         string `json:"test_type,omitempty"`
	Quantity         uint   `json:"quantity,omitempty"`
	Description      string `json:"description,omitempty"`
}
