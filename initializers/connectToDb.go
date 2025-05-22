package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dbString := os.Getenv("DB_STRING")
	DB, err = gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		panic("Failed to connec to DB")
	}
}
