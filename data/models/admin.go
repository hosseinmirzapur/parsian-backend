package models

import "github.com/hosseinmirzapur/parsian-backend/common"

type Admin struct {
	BaseModel
	Username string `gorm:"type:string;size:20;not null;unique"`
	Name     string `gorm:"type:string;size:50;null"`
	Password string
	Role     common.AdminRole `gorm:"type:string;size:20;not null"`
}
