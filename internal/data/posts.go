package data

import (
	"time"

	"github.com/bartvanbenthem/gofound-blogger/internal/validator"
)

type Post struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author,omitempty"`
	ImgURLs   []string  `json:"img_urls,omitempty"`
}

func ValidatePost(v *validator.Validator, post *Post) {
	v.Check(post.Title != "", "title", "must be provided")
	v.Check(len(post.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(post.Content != "", "content", "must be provided")
	v.Check(len(post.Title) <= 10000, "content", "must not be more than 10000 bytes long")
	v.Check(validator.Unique(post.ImgURLs), "img_urls", "must not contain duplicate values")
}
