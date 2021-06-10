package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dunkbing/sfw-checker-viet/backend/database"
	"github.com/dunkbing/sfw-checker-viet/backend/models"
	"github.com/dunkbing/sfw-checker-viet/backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func Register(registerUser models.RegisterUser) (models.User, error) {
	var dbUser models.User

	if utils.ValidEmail(registerUser.Email) {
		database.DB.Where("email = ?", registerUser.Email).First(&dbUser)
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

	database.DB.Create(&dbUser)

	return dbUser, nil
}

func Login(loginUser models.LoginUser) (string, error, int) {
	var dbUser models.User

	if err := database.DB.Where("email = ?", loginUser.Email).First(&dbUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("user not found"), http.StatusNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
		return "", errors.New("incorrect password"), http.StatusBadRequest
	}

	claims := jwt.StandardClaims{
		Id:        string(int32(dbUser.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.FormatInt(int64(dbUser.Id), 10),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	return token, err, 0
}

func User(id string) (models.User, error, int) {
	var user models.User
	res := database.DB.Where("id = ?", id).First(&user)
	err := res.Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, err, http.StatusNotFound
	}

	return user, err, 0
}
