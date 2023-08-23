package dto

type CreateOrderRequest struct {
	PhoneNumber  string `json:"phone_number" validate:"required,numeric,size=11"`
	CustomerName string `json:"customer_name" validate:"required"`
}

type UpdateOrderRequest struct {
	PhoneNumber  string `json:"phone_number,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
}
