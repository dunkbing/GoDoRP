package routes

import (
	"github.com/dunkbing/sfw-checker-viet/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("login", controllers.Login)
	app.Post("logout", controllers.Logout)
	app.Get("user", controllers.User)
}
