package controllers

import (
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User = models.User{FirstName: "John", LastName: "Doe", Email: "rewr@rer.c", Password: "pass"}
	return c.JSON(user)
}
