package handlers

import (
	"net/http"

	"github.com/dunkbing/sfw-checker-viet/backend/data"
)

func GetAll(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	posts := data.GetPosts()

	err := data.ToJson(posts, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
