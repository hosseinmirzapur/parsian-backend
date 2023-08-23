package dto

type CreateOrderItemRequest struct {
	Name             string `json:"name" validation:"required"`
	AllowSandPaper   bool   `json:"allow_sand_paper" validation:"required"`
	AllowDestruction bool   `json:"allow_destruction" validation:"required"`
	Status           string `json:"status" validation:"required"`
	TestType         string `json:"test_type" validation:"required"`
	Quantity         uint   `json:"quantity" validation:"required,min=1"`
	Description      string `json:"description"`
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
