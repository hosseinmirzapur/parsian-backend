package models

import "time"

type BaseModel struct {
	Id uint `gorm:"primary_key;auto_increment"`

	CreatedAt time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
	UpdatedAt time.Time `gorm:"type:TIMESTAMP with time zone;null"`
}
