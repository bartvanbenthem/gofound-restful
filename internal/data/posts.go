package data

import "time"

type Post struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author,omitempty"`
}
