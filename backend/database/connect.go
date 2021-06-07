package database

import (
	"fmt"
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=sa dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("couldn't connect to the database")
	}

	DB = db

	err = db.AutoMigrate(&models.User{}, &models.PasswordReset{})
	if err != nil {
		panic(fmt.Sprintf("Error migrating db: %s", err.Error()))
		return
	}
}

func Close() {
	sqlDb, err := DB.DB()
	if err != nil {
		panic(err.Error())
	}
	_ = sqlDb.Close()
}
