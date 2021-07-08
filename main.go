package main

import (
	"github.com/dunkbing/sfw-checker-viet/backend/api"
	database "github.com/dunkbing/sfw-checker-viet/backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// database.Connect()
	defer database.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	server := api.New()
	server.Init()
}
