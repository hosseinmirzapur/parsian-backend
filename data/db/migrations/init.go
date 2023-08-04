package migrations

import (
	"log"

	"github.com/hosseinmirzapur/parsian-backend/data/db"
	"gorm.io/gorm"
)

func Up() {
	database := db.GetDB()
	createTables(database)
}

func Down() {

}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// TODO: add models to the tables array

	// Migrate
	err := database.Migrator().CreateTable(tables...)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("tables created")
}

/**
Uncomment this...
*/

// func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
// 	if !database.Migrator().HasTable(model) {
// 		tables = append(tables, model)
// 	}
// 	return tables
// }
