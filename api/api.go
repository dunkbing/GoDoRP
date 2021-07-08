package api

import (
	"fmt"
	"net/http"
	"os"

	service "github.com/dunkbing/sfw-checker-viet/backend/services"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	ApiRoot fiber.Router
	Auth    fiber.Router
}

type Api struct {
	app        *fiber.App
	BaseRoutes *Routes
}

func New() *Api{
	return &Api{}
}

const PrefixApi = "api"

func StatusOk(c *fiber.Ctx, json interface{}) error {
	c.Status(http.StatusOK)
	return c.JSON(fiber.Map{
		"status": "success",
		"result": json,
	})
}

func StatusError(c *fiber.Ctx, httpError service.HttpError) error {
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

func (api *Api) Init() {
	app := fiber.New()
	routes := &Routes{}

	api.app = app
	api.BaseRoutes = routes

	app.Static("/", "./frontend")

	api.BaseRoutes.ApiRoot = app.Group(PrefixApi)
	api.BaseRoutes.Auth = api.BaseRoutes.ApiRoot.Group("auth")

	api.InitAuth()
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
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
