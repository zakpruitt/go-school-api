package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"school-api/models"
)

const DSN = "host=localhost port=5432 user=postgres dbname='School API DB' sslmode=disable password=User123"

func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database models:", err)
	}

	return db
}
