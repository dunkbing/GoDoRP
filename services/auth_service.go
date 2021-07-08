package service

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"github.com/dunkbing/sfw-checker-viet/backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {}

func (authService) Register(registerUser models.RegisterUser) (models.User, error) {
	var dbUser models.User

	if utils.ValidEmail(registerUser.Email) {
		Database.Where("email = ?", registerUser.Email).First(&dbUser)
		if dbUser.Id != 0 {
			return dbUser, errors.New("email already in use")
		}
	} else {
		return dbUser, errors.New("invalid email")
	}

	if !utils.ValidPassword(registerUser.Password) {
		return dbUser, errors.New("invalid password")
	}

	if registerUser.Password != registerUser.ConfirmPass {
		return dbUser, errors.New("passwords do not match")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 14)
	dbUser.FirstName = registerUser.FirstName
	dbUser.LastName = registerUser.LastName
	dbUser.Email = registerUser.Email
	dbUser.Password = string(password)

	err := Database.Create(&dbUser).Error

	return dbUser, err
}

func (authService) Login(loginUser models.LoginUser) (string, *HttpError) {
	var dbUser models.User

	if err := Database.Where("email = ?", loginUser.Email).First(&dbUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return "", &HttpError{
			Message: "user not found",
			StatusCode: http.StatusNotFound,
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
		return "", &HttpError{
			Message: "incorrect password",
			StatusCode: http.StatusBadRequest,
		}
	}

	claims := jwt.StandardClaims{
		Id:        string(int32(dbUser.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.FormatInt(int64(dbUser.Id), 10),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	return token, &HttpError{
		Message: err.Error(),
	}
}

func (authService) User(id string) (models.User, *HttpError) {
	var user models.User
	res := Database.Where("id = ?", id).First(&user)
	err := res.Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, &HttpError{
			Message: "User not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return user, &HttpError{
		Message: err.Error(),
	}
}

func GetAuthService() authService {
	return authService{}
}