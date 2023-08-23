package services

import (
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/common"
	"github.com/hosseinmirzapur/parsian-backend/data/db"
	"github.com/hosseinmirzapur/parsian-backend/data/models"
)

func CreateOrderItem(data *dto.CreateOrderItemRequest, filepath string, orderId int) (models.OrderItem, error) {
	dbClient := db.GetDB()
	order := models.Order{
		BaseModel: models.BaseModel{
			Id: uint(orderId),
		},
	}
	orderItem := models.OrderItem{}
	err := dbClient.First(&order).Error
	if err != nil {
		return orderItem, err
	}

	orderItem.AllowDestruction = data.AllowDestruction
	orderItem.AllowSandPaper = data.AllowSandPaper
	orderItem.Description = data.Description
	orderItem.Quantity = data.Quantity
	orderItem.Name = data.Name
	orderItem.Status = common.OrderStatus(data.Status)
	orderItem.TestType = common.TestType(data.TestType)
	orderItem.FilePath = filepath
	orderItem.Order = order

	err = dbClient.Create(&orderItem).Error
	return orderItem, err
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
