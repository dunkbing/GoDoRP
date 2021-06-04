package data

import (
	"net/http"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Posts []*Post

func GetAll() []Post {
	posts := []Post{}
	DB.Find(&posts)
	return posts
}

func Create(req *http.Request) {
	// fetch post from context
	post := req.Context().Value("").(Post)

	DB.Create(&post)
}
