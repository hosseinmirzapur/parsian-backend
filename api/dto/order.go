package dto

import "github.com/hosseinmirzapur/parsian-backend/common"

type CreateOrderRequest struct {
	PhoneNumber  string `json:"phone_number" validation:"required"`
	CustomerName string `json:"customer_name" validation:"required"`
}

type UpdateOrderRequest struct {
	PhoneNumber  string `json:"phone_number,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
}

type ChangeOrderStatusRequest struct {
	Status common.OrderStatus `json:"status" validation:"required"`
}
