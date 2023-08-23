package models

type Order struct {
	BaseModel
	OrderItems   *[]OrderItem
	PhoneNumber  string `gorm:"type:string;size:11;not_null"`
	CustomerName string `gorm:"type:string;not_null"`
	SpecialId    string `gorm:"type:string;size:8;not_null"`
}

type OrderItem struct {
	BaseModel
	Name             string `gorm:"type:string;size:20;not_null"`
	AllowSandPaper   bool   `gorm:"default:false"`
	AllowDestruction bool   `gorm:"default:false"`
	Status           string `gorm:"default:pending"`
	TestType         string `gorm:"default:analyze"`
	Quantity         uint   `gorm:"default:1"`
	Description      string `gorm:"type:text"`
	FilePath         string `gorm:"type:string;size:255"`
	Order            Order  `gorm:"foreignkey:OrderId;constraint:onDelete:Cascade"`
	OrderId          uint
}
