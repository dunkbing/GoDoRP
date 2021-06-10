package models

type User struct {
	Id        uint   `json:"id" swaggerignore:"true"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-" swaggerignore:"true"`
}

type RegisterUser struct {
	FirstName   string `json:"firstName" example:"Bing"`
	LastName    string `json:"lastName" example:"Bui"`
	Email       string `json:"email" example:"bing@dep.chai"`
	Password    string `json:"password" example:"Bingvip69."`
	ConfirmPass string `json:"confirmPass" example:"Bingvip69."`
}

type LoginUser struct {
	Email    string `json:"email" example:"bing@dep.chai"`
	Password string `json:"password" example:"Bingvip69."`
}

type LoginResponse struct {
	Jwt string `json:"jwt"`
}
