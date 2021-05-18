package data

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() {
	USER := os.Getenv("POSTGRES_USER")
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("POSTGRES_DB")

	dns := fmt.Sprintf("host=postgres port=%v user=%v database=%v password=%v sslmode=disable", PORT, USER, DB_NAME, PASSWORD)
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err == nil {
			println("connected to db")
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate()
}
