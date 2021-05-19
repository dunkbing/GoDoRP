package data

import "gorm.io/gorm"

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
