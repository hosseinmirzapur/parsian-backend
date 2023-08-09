package services

import (
	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/api/helper"
	"github.com/hosseinmirzapur/parsian-backend/common"
	"github.com/hosseinmirzapur/parsian-backend/data/db"
	"github.com/hosseinmirzapur/parsian-backend/data/models"
)

func AllOrders() ([]models.Order, error) {
	dbClient := db.GetDB()
	var orders []models.Order
	err := dbClient.Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func GetOrderById(id int) (models.Order, error) {
	order := models.Order{
		BaseModel: models.BaseModel{
			Id: uint(id),
		},
	}
	dbClient := db.GetDB()
	err := dbClient.Preload("OrderItems").Find(&order).Error
	return order, err
}

func CreateOrder(data *dto.CreateOrderRequest) (models.Order, error) {
	dbClient := db.GetDB()
	order := models.Order{}
	order.SpecialId = helper.GenerateOrderCode()
	order.CustomerName = data.CustomerName
	order.PhoneNumber = data.PhoneNumber
	err := dbClient.Create(&order).Error
	return order, err
}

func UpdateOrder(data *dto.UpdateOrderRequest, id int) (models.Order, error) {
	dbClient := db.GetDB()
	order := models.Order{
		BaseModel: models.BaseModel{
			Id: uint(id),
		},
	}
	err := dbClient.Model(&order).Preload("OrderItems").Updates(data).Error
	return order, err
}

func DeleteOrder(id int) error {
	dbClient := db.GetDB()
	return dbClient.Delete(&models.Order{BaseModel: models.BaseModel{Id: uint(id)}}).Error
}

func ChangeOrderStatus(id int, status common.OrderStatus) (models.Order, error) {
	var err error
	dbClient := db.GetDB()
	order := models.Order{
		BaseModel: models.BaseModel{
			Id: uint(id),
		},
	}
	for _, item := range *order.OrderItems {
		err = dbClient.Model(&item).Update("status", status).Error
		if err != nil {
			return order, err
		}
	}
	return order, nil
}
