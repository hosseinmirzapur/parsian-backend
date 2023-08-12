package dto

type CreateOrderRequest struct {
	PhoneNumber  string `json:"phone_number" validation:"required"`
	CustomerName string `json:"customer_name" validation:"required"`
}

type UpdateOrderRequest struct {
	PhoneNumber  string `json:"phone_number,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
}
