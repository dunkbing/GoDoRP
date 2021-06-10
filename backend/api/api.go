package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/dunkbing/sfw-checker-viet/backend/docs"
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

const PrefixApi = "api"

func StatusOk(c *fiber.Ctx, json interface{}) error {
	c.Status(http.StatusOK)
	return c.JSON(json)
}

func StatusCreated(c *fiber.Ctx, json interface{}) error {
	c.Status(http.StatusCreated)
	return c.JSON(json)
}

func StatusBadRequest(c *fiber.Ctx, appError HttpError) error {
	c.Status(http.StatusBadRequest)
	return c.JSON(appError)
}

func StatusNotFound(c *fiber.Ctx, appError HttpError) error {
	c.Status(http.StatusNotFound)
	return c.JSON(appError)
}

// Init
// @BasePath /api
func Init(app *fiber.App) {
	api := &API{
		app:        app,
		BaseRoutes: &Routes{},
	}

	app.Get("/swagger/*", swagger.Handler)

	api.BaseRoutes.ApiRoot = app.Group(PrefixApi)
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
