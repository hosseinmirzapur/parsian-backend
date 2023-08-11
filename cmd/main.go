package main

import (
	"log"

	"github.com/hosseinmirzapur/parsian-backend/api"
	"github.com/hosseinmirzapur/parsian-backend/config"
	"github.com/hosseinmirzapur/parsian-backend/data/db"
	"github.com/hosseinmirzapur/parsian-backend/data/db/migrations"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	cfg := config.NewConfig("development")

	if cfg.IsDevelopment() {
		err = godotenv.Load("./.env")

		if err != nil {
			log.Fatal("Error loading env file")
		}
	}

	err = db.InitDB()
	defer db.CloseDB()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	migrations.Up()

	api.InitServer()
}
