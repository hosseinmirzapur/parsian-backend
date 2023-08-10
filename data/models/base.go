package models

import "time"

type BaseModel struct {
	Id uint `gorm:"primaryKey;autoIncrement"`

	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"null"`
}
