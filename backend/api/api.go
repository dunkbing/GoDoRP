package api

import (
	"net/http"

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

func StatusOk(c *fiber.Ctx, json interface{}) error {
	c.Status(http.StatusOK)
	return c.JSON(json)
}

func StatusCreated(c *fiber.Ctx, json interface{}) error {
	c.Status(http.StatusCreated)
	return c.JSON(json)
}

func StatusBadRequest(c *fiber.Ctx, appError AppError) error {
	c.Status(http.StatusBadRequest)
	return c.JSON(appError)
}

func StatusNotFound(c *fiber.Ctx, appError AppError) error {
	c.Status(http.StatusNotFound)
	return c.JSON(appError)
}

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
