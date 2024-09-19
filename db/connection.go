package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DSN = "host=digicel-dev-flex.postgres.database.azure.com user=xxepin password=migracion dbname=testing port=5432"
)

var DB *gorm.DB

func DBconnction() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal("🤒")
	} else {
		log.Println("Db connection established")
	}

}
