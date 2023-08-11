package dto

import "github.com/hosseinmirzapur/parsian-backend/common"

type CreateAdminRequest struct {
	Username string           `json:"username" validation:"required"`
	Name     string           `json:"name" validation:"required"`
	Password string           `json:"password" validation:"required"`
	Role     common.AdminRole `json:"role" validation:"required"`
}

type UpdateAdminRequest struct {
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
}
