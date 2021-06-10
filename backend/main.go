package main

import (
	"github.com/dunkbing/sfw-checker-viet/backend/api"
	"github.com/dunkbing/sfw-checker-viet/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"os"
)

// @title Learning go
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath
// @schemes http
func main() {
	_ = godotenv.Load()

	database.Connect()

	defer database.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     os.Getenv("ALLOW_ORIGIN"),
	}))

	api.Init(app)

	_ = app.Listen(":8080")
}
