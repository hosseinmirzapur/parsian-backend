package seeders

import (
	"log"
	"os"

	"github.com/hosseinmirzapur/parsian-backend/data/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(dbClient *gorm.DB) {
	admins := []models.Admin{}

	dbClient.Find(&admins)

	if len(admins) == 0 {
		err := seedAdmins(dbClient)

		if err != nil {
			log.Fatal("Unable to seed the database")
		}
	}

}

func seedAdmins(dbClient *gorm.DB) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("DEFAULT_ADMIN_PASSWORD")), 10)

	admin := models.Admin{
		Username: os.Getenv("DEFAULT_ADMIN_USERNAME"),
		Name:     os.Getenv("DEFAULT_ADMIN_NAME"),
		Role:     "admin",
		Password: string(password),
	}

	return dbClient.Create(&admin).Error
}
