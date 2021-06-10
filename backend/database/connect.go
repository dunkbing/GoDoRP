package database

import (
	"fmt"
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

func Connect() {
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbPort   = os.Getenv("DB_PORT")
		dbName   = os.Getenv("DB_NAME")
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
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
