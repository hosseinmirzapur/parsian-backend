package dto

type LoginRequest struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"password" validation:"required,min=8"`
}
