package services

import (
	"fmt"

	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/api/helper"
	"github.com/hosseinmirzapur/parsian-backend/data/db"
	"github.com/hosseinmirzapur/parsian-backend/data/models"
	"github.com/hosseinmirzapur/parsian-backend/utils"
)

func AllOrders() ([]models.Order, error) {
	dbClient := db.GetDB()
	var orders []models.Order
	query := dbClient.Preload("OrderItems").Find(&orders).Limit(50)

	return orders, query.Error
}

func GetOrderById(id int) (models.Order, error) {
	order := models.Order{
		BaseModel: models.BaseModel{
			Id: uint(id),
		},
	}
	dbClient := db.GetDB()
	err := dbClient.Preload("OrderItems").First(&order).Error
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
	err := dbClient.First(&order).Preload("OrderItems").Updates(data).Error
	return order, err
}

func DeleteOrder(id int) error {
	dbClient := db.GetDB()
	order := models.Order{
		BaseModel: models.BaseModel{
			Id: uint(id),
		},
	}
	res := dbClient.Delete(&order)

	if res.RowsAffected == 0 {
		return fmt.Errorf("order with id %d not found", id)
	}
	return res.Error
}

func FindOrderBySpecialId(specialId string) (models.Order, error) {
	dbClient := db.GetDB()

	order := models.Order{}

	err := dbClient.Where("special_id = ?", specialId).Preload("OrderItems").First(&order).Error

	return order, err
}

func GetExcelFile() ([]byte, error) {

	// Retrieve orders
	dbClient := db.GetDB()
	var orderItems []models.OrderItem
	err := dbClient.Find(&orderItems).Error
	if err != nil {
		return nil, err
	}
	// Exports data into export.xlsx
	excel, err := utils.ExcelExport(orderItems)
	if err != nil {
		return nil, err
	}

	bufferArray, err := excel.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	// Convetting file to buffer
	return bufferArray.Bytes(), nil

}
