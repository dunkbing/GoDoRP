package api

import (
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	ApiRoot fiber.Router
	Auth    fiber.Router
}

type API struct {
	app        *fiber.App
	BaseRoutes *Routes
}

const API_PREFIX = "api"

func Init(app *fiber.App) {
	api := &API{
		app:        app,
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.ApiRoot = app.Group(API_PREFIX)
	api.BaseRoutes.Auth = api.BaseRoutes.ApiRoot.Group("auth")

	api.InitAuth()
}

// func Setup(app *fiber.App) {
// 	api := app.Group("api")
// 	auth := api.Group("auth")
// 	auth.Post("register", controllers.Register)
// 	auth.Post("login", controllers.Login)
// 	auth.Post("logout", controllers.Logout)
// 	auth.Get("user", controllers.User)
// 	auth.Post("forgot", controllers.Forgot)
// 	auth.Post("reset", controllers.Reset)
// }
