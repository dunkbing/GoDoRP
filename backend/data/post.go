package data

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Posts []*Post

func GetPosts() Posts {
	return posts
}

// dummy data
var posts = Posts{
	&Post{
		ID:          1,
		Title:       "post 1",
		Description: "des 1",
	},
	&Post{
		ID:          2,
		Title:       "post 2",
		Description: "des 2",
	},
}
