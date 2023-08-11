package models

import "github.com/hosseinmirzapur/parsian-backend/common"

type Admin struct {
	BaseModel
	Username string           `gorm:"type:string;size:20;not_null;unique"`
	Name     string           `gorm:"type:string;size:50;null"`
	Role     common.AdminRole `gorm:"type:string;size:20;not_null"`
	Password string
}
