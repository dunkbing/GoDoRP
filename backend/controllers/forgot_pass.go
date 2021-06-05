package controllers

import (
	"math/rand"
	"net/http"
	"net/smtp"

	"github.com/dunkbing/sfw-checker-viet/backend/database"
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"github.com/dunkbing/sfw-checker-viet/backend/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	email := data["email"]

	if !utils.ValidEmail(email) {
		c.Status(http.StatusBadRequest)
		return c.JSON(utils.AppError{
			Message:    "invalid email",
			StatusCode: http.StatusBadRequest,
		})
	}

	var user models.User

	res := database.DB.Model(&user).Where("email = ?", email).First(&user)

	if res.Error != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(fiber.Map{
			"message": "some error occur",
		})
	}

	if user.Email == "" {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "email does not exist",
		})
	}

	token := RandStringRunes(12)

	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	database.DB.Create(&passwordReset)

	from := "admin@dunkbing.com"
	to := []string{
		data["email"],
	}

	url := "http://localhost:3000/reset/" + token

	message := []byte("Click <a href=\"" + url + "\">here</a> to reset your password!")

	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, message)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Reset(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["confirm_pass"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	var passwordReset = models.PasswordReset{}

	if res := database.DB.Where("token = ?", data["token"]).Last(&passwordReset); res.Error != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid token!",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
