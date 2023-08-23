package services

import (
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/data/db"
	"github.com/hosseinmirzapur/parsian-backend/data/models"
)

func CreateOrderItem(data *dto.CreateOrderItemRequest, filepath string, orderId int) error {
	dbClient := db.GetDB()
	order := models.Order{
		BaseModel: models.BaseModel{
			Id: uint(orderId),
		},
	}

	err := dbClient.First(&order).Error
	if err != nil {
		return err
	}

	err = dbClient.Create(&models.OrderItem{
		AllowDestruction: data.AllowDestruction == 1,
		AllowSandPaper:   data.AllowSandPaper == 1,
		Description:      data.Description,
		Quantity:         data.Quantity,
		Name:             data.Name,
		Status:           data.Status,
		TestType:         data.TestType,
		FilePath:         filepath,
		Order:            order,
	}).Error
	return err
}

func UpdateOrderItem(data *dto.UpdateOrderItemRequest, oiId int) (models.OrderItem, error) {
	dbClient := db.GetDB()
	orderItem := models.OrderItem{
		BaseModel: models.BaseModel{
			Id: uint(oiId),
		},
	}
	err := dbClient.First(&orderItem).Updates(data).Error
	return orderItem, err
}

func DeleteOrderItem(oiId int) error {
	dbClient := db.GetDB()
	orderItem := models.OrderItem{
		BaseModel: models.BaseModel{
			Id: uint(oiId),
		},
	}

	err := dbClient.First(&orderItem).Error
	if err != nil {
		return err
	}

	return dbClient.Delete(&orderItem).Error
}
