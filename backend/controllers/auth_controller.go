package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dunkbing/sfw-checker-viet/backend/database"
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type claims struct {
	Id int
	jwt.StandardClaims
}

func Register(c *fiber.Ctx) error {
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
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	var user models.User
	user.FirstName = data["first_name"]
	user.LastName = data["last_name"]
	user.Email = data["email"]
	user.Password = string(password)

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "password not correct",
		})
	}

	claims := jwt.StandardClaims{
		Id:        string(int32(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.FormatInt(int64(user.ID), 10),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"jwt": token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	claims := jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(cookie, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	fmt.Println(claims)

	if err != nil || !token.Valid {
		c.Status(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// fmt.Println(token.Claims)
	id := claims.Issuer
	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}
