package main

import (
	"github.com/dunkbing/sfw-checker-viet/backend/api"
	"github.com/dunkbing/sfw-checker-viet/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	sqlDb, err := database.DB.DB()
	if err != nil {
		panic(err.Error())
	}
	defer sqlDb.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	api.Init(app)

	app.Listen(":8080")
}
