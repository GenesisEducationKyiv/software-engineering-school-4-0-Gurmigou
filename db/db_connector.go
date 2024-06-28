package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"se-school-case/pkg/constants"
)

var DB *gorm.DB

func ConnectToDatabase() *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(constants.DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	return DB
}
