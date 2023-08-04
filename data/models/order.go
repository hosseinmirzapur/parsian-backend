package models

import "github.com/hosseinmirzapur/parsian-backend/common"

type Order struct {
	BaseModel
	OrderItems   *[]OrderItem
	PhoneNumber  string `gorm:"type:string;size:11;not null"`
	CustomerName string `gorm:"type:string;size:255;not null"`
	SpecialId    string `gorm:"type:string;size:8;not null"`
}

type OrderItem struct {
	BaseModel
	Name             string             `gorm:"type:string;size:20;not null"`
	AllowSandPaper   bool               `gorm:"default:false"`
	AllowDestruction bool               `gorm:"default:false"`
	Order            Order              `gorm:"foreignkey:OrderId;constraint:onDelete:Cascade"`
	Status           common.OrderStatus `gorm:"default:pending"`
	TestType         common.TestType    `gorm:"default:analyze"`
	Quantity         uint               `gorm:"default:1"`
	Description      string             `gorm:"type:text"`
	FileTitle        string             `gorm:"type:string;size:255"`
	FileType         string             `gorm:"type:string;size:255"`
	OrderId          uint
}
