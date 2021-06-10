package api

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"net/http"
	"net/smtp"

	"github.com/dgrijalva/jwt-go"
	"github.com/dunkbing/sfw-checker-viet/backend/database"
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"github.com/dunkbing/sfw-checker-viet/backend/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Summary Register a new user.
// @Description Register a new user.
// @Tags register
// @Accept json
// @Produce json
// @Param user body models.RegisterUser true "register user"
// @Success 201 {object} models.User
// @Failure 400 {object} AppError
// @Router /api/auth/register [post]
func register(c *fiber.Ctx) error {

	var registerUser models.RegisterUser
	if err := c.BodyParser(&registerUser); err != nil {
		return StatusBadRequest(c, AppError{Message: err.Error()})
	}

	var user models.User

	if utils.ValidEmail(registerUser.Email) {
		database.DB.Where("email = ?", registerUser.Email).First(&user)
		if user.Id != 0 {
			return StatusBadRequest(c, AppError{
				Message: "Email already in use",
			})
		}
	} else {
		return StatusBadRequest(c, AppError{
			Message: "Invalid email",
		})
	}

	if !utils.ValidPassword(registerUser.Password) {
		return StatusBadRequest(c, AppError{
			Message: "invalid password",
		})
	}

	if registerUser.Password != registerUser.ConfirmPass {
		return StatusBadRequest(c, AppError{
			Message: "Passwords do not match",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 14)
	user.FirstName = registerUser.FirstName
	user.LastName = registerUser.LastName
	user.Email = registerUser.Email
	user.Password = string(password)

	database.DB.Create(&user)

	return StatusCreated(c, user)
}

// @Summary Login to user
// @Description login.
// @Tags login
// @Accept json
// @Produce json
// @Param user body models.User true "login user"
// @Success 201 {object} models.User
// @Failure 400 {object} interface{}
// @Router /api/auth/login [post]
func login(c *fiber.Ctx) error {
	var loginUser models.LoginUser
	if err := c.BodyParser(&loginUser); err != nil {
		return StatusBadRequest(c, AppError{
			Message: err.Error(),
		})
	}

	var dbUser models.User

	if err := database.DB.Where("email = ?", loginUser.Email).First(&dbUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return StatusNotFound(c, AppError{
			Message: "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
		return StatusBadRequest(c, AppError{
			Message: "Incorrect password",
		})
	}

	claims := jwt.StandardClaims{
		Id:        string(int32(dbUser.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.FormatInt(int64(dbUser.Id), 10),
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

// ShowUser godoc
// @Summary Show current user
// @Description get current user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} models.User
// @Failure 400 {object} AppError
// @Router /api/auth/user [get]
func user(c *fiber.Ctx) error {
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

	return StatusOk(c, user)
}

func logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func forgot(c *fiber.Ctx) error {
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
		return c.JSON(AppError{
			Message:    "invalid email",
			StatusCode: http.StatusBadRequest,
		})
	}

	var user models.User

	res := database.DB.Model(&user).Where("email = ?", email).First(&user)

	if res.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "some error occur",
		})
	}

	if user.Email == "" {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "email does not exist",
		})
	}

	token := utils.RandStringRunes(12)

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

func reset(c *fiber.Ctx) error {
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

func (api *API) InitAuth() {
	auth := api.BaseRoutes.Auth
	auth.Post("register", register)
	auth.Post("login", login)
	auth.Post("logout", logout)
	auth.Get("user", user)
	auth.Post("forgot", forgot)
	auth.Post("reset", reset)
}
