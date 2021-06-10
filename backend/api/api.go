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
	return c.JSON(fiber.Map{
		"status": "success",
		"result": json,
	})
}

func StatusError(c *fiber.Ctx, httpError HttpError) error {
	c.Status(httpError.StatusCode)
	return c.JSON(fiber.Map{
		"status": "failed",
		"result": httpError,
	})
}

func StatusCreated(c *fiber.Ctx, json interface{}) error {
	c.Status(http.StatusCreated)
	return c.JSON(fiber.Map{
		"status": "success",
		"result": json,
	})
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
