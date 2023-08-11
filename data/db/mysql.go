package db

import (
	"fmt"
	"log"
	"os"

	"github.com/hosseinmirzapur/parsian-backend/data/db/seeders"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDB() error {
	var err error

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	dbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()

	if err != nil {
		return err
	}

	log.Println("Database connection established...")

	seeders.Seed(dbClient)

	log.Println("Database seeded...")
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	con, _ := dbClient.DB()
	con.Close()
}
