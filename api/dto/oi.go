package dto

type CreateOrderItemRequest struct {
	Name             string `json:"name" validation:"required"`
	AllowSandPaper   bool   `json:"allow_sand_paper" validation:"required,boolean"`
	AllowDestruction bool   `json:"allow_destruction" validation:"required,boolean"`
	Status           string `json:"status" validation:"required,oneof=pending partial office paid"`
	TestType         string `json:"test_type" validation:"required,oneof=analyze hardness both"`
	Quantity         uint   `json:"quantity" validation:"required,min=1"`
	Description      string `json:"description"`
}

type UpdateOrderItemRequest struct {
	Name             string `json:"name,omitempty"`
	AllowSandPaper   bool   `json:"allow_sand_paper,omitempty" validation:"boolean"`
	AllowDestruction bool   `json:"allow_destruction,omitempty" validation:"boolean"`
	Status           string `json:"status,omitempty" validation:"oneof=pending partial office paid"`
	TestType         string `json:"test_type,omitempty" validation:"oneof=analyze hardness both"`
	Quantity         uint   `json:"quantity,omitempty"`
	Description      string `json:"description,omitempty"`
}
