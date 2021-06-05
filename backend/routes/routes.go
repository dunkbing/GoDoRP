package routes

import (
	"github.com/dunkbing/sfw-checker-viet/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)
	app.Post("api/logout", controllers.Logout)
	app.Get("api/user", controllers.User)
	app.Post("api/forgot", controllers.Forgot)
	app.Post("api/reset", controllers.Reset)
}
