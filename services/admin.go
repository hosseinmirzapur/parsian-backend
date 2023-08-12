package services

import (
	"fmt"

	"github.com/hosseinmirzapur/parsian-backend/api/dto"
	"github.com/hosseinmirzapur/parsian-backend/common"
	"github.com/hosseinmirzapur/parsian-backend/data/db"
	"github.com/hosseinmirzapur/parsian-backend/data/models"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(username, password string) (string, models.Admin, error) {

	var err error
	admin := models.Admin{
		Username: username,
	}

	dbClient := db.GetDB()

	err = dbClient.Where("username = ?", username).First(&admin).Error

	if err != nil {
		return "", admin, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))

	if err != nil {
		return "", admin, err
	}

	token, err := GenerateJWTToken(fmt.Sprint(admin.BaseModel.Id))

	return token, admin, err
}

func CreateAdmin(username, password, name string, role common.AdminRole) (models.Admin, error) {
	dbClient := db.GetDB()

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if role == "admin" {
		return models.Admin{}, fmt.Errorf("admin can only create employees")
	}

	admin := models.Admin{
		Username: username,
		Name:     name,
		Password: string(passwordHash),
		Role:     role,
	}

	err := dbClient.Create(&admin).Error

	return admin, err
}

func UpdateAdmin(data *dto.UpdateAdminRequest, id uint) (models.Admin, error) {
	dbClient := db.GetDB()

	admin := models.Admin{
		BaseModel: models.BaseModel{
			Id: id,
		},
	}

	err := dbClient.First(&admin).Updates(data).Error

	return admin, err
}
func DeleteAdmin(id uint) error {
	dbClient := db.GetDB()

	admin := models.Admin{
		BaseModel: models.BaseModel{
			Id: id,
		},
	}

	res := dbClient.Delete(&admin)

	if res.RowsAffected == 0 {
		return fmt.Errorf("admin not found")
	}

	return res.Error
}
