package api

import (
	"fmt"
	"time"

	"net/http"
	"net/smtp"

	"github.com/dgrijalva/jwt-go"
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	service "github.com/dunkbing/sfw-checker-viet/backend/services"
	"github.com/dunkbing/sfw-checker-viet/backend/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


// @Summary Register a new user.
// @Description Register a new user.
// @Tags register
// @Accept json
// @Produce json
// @Param user body models.RegisterUser true "register user"
// @Success 201 {object} models.User
// @Failure 400 {object} HttpError
// @Router /api/auth/register [post]
func register(c *fiber.Ctx) error {
	var registerUser models.RegisterUser
	if err := c.BodyParser(&registerUser); err != nil {
		return StatusError(c, service.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}
	var authService = service.GetAuthService()

	dbUser, err := authService.Register(registerUser)

	if err != nil {
		return StatusError(c, service.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}

	return StatusCreated(c, dbUser)
}

// @Summary Login to user
// @Description login.
// @Tags login
// @Accept json
// @Produce json
// @Param user body models.LoginUser true "login user"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} interface{}
// @Router /api/auth/login [post]
func login(c *fiber.Ctx) error {
	var loginUser models.LoginUser
	if err := c.BodyParser(&loginUser); err != nil {
		return StatusError(c, service.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}
	var authService = service.GetAuthService()

	token, err := authService.Login(loginUser)

	if err != nil {
		return StatusError(c, *err)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return StatusOk(c, models.LoginResponse{Jwt: token})
}

// ShowUser godoc
// @Summary Show current user
// @Description get current user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.User
// @Failure 400 {object} HttpError
// @Router /api/auth/user [get]
func user(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	claims := jwt.StandardClaims{}

	fmt.Println(claims)
	

	if token, err := jwt.ParseWithClaims(cookie, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	}); err != nil || !token.Valid {
		c.Status(http.StatusUnauthorized)
		return StatusError(c, service.HttpError{
			Message:    "unauthenticated",
			StatusCode: http.StatusForbidden,
		})
	}

	var authService = service.GetAuthService()

	// fmt.Println(token.Claims)
	id := claims.Issuer
	user, err := authService.User(id)
	if err != nil {
		return StatusError(c, *err)
	}

	return StatusOk(c, user)
}

// @Summary Logout
// @Tags logout
// @Produce json
// @Success 200 {object} interface{}
// @Router /api/auth/logout [post]
func logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return StatusOk(c, fiber.Map{
		"message": "success",
	})
}

// @Summary Forgot password
// @Tags forgot
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 500 {object} HttpError
// @Failure 400 {object} HttpError
// @Failure 404 {object} HttpError
// @Router /api/auth/forgot [post]
func forgot(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return StatusError(c, service.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
	}

	email := data["email"]

	if !utils.ValidEmail(email) {
		return StatusError(c, service.HttpError{
			Message:    "invalid email",
			StatusCode: http.StatusBadRequest,
		})
	}

	var user models.User

	res := service.Database.Model(&user).Where("email = ?", email).First(&user)

	if res.Error != nil {
		return StatusError(c, service.HttpError{
			Message:    "some error occur",
			StatusCode: http.StatusInternalServerError,
		})
	}

	if user.Email == "" {
		return StatusError(c, service.HttpError{
			Message:    "email does not exist",
			StatusCode: http.StatusNotFound,
		})
	}

	token := utils.RandStringRunes(12)

	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	service.Database.Create(&passwordReset)

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

// @Summary Reset password
// @Tags reset
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 400 {object} HttpError
// @Router /api/auth/reset [post]
func reset(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["confirm_pass"] {
		return StatusError(c, service.HttpError{
			Message:    "Passwords do not match",
			StatusCode: http.StatusBadRequest,
		})
	}

	var passwordReset = models.PasswordReset{}

	if res := service.Database.Where("token = ?", data["token"]).Last(&passwordReset); res.Error != nil {
		return StatusError(c, service.HttpError{
			Message:    "Invalid token!",
			StatusCode: http.StatusBadRequest,
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	service.Database.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return StatusOk(c, fiber.Map{
		"message": "success",
	})
}

func (api *Api) InitAuth() {
	auth := api.BaseRoutes.Auth
	auth.Post("register", register)
	auth.Post("login", login)
	auth.Post("logout", logout)
	auth.Get("user", user)
	auth.Post("forgot", forgot)
	auth.Post("reset", reset)
}
