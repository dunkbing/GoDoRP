package main

import (
	"github.com/dunkbing/sfw-checker-viet/backend/database"
	"github.com/dunkbing/sfw-checker-viet/backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}
