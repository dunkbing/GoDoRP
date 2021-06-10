package models

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
}

type RegisterUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirmPass"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
