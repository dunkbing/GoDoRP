package database

import (
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=sa dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to the database")
	}

	DB = db

	db.AutoMigrate(&models.User{})
}
