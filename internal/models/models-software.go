package models

import "time"

// Software is the type for movies
type Software struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Year             int            `json:"year"`
	ReleaseDate      time.Time      `json:"release_date"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	SoftwareCategory map[int]string `json:"categories"`
}

// Category is the type for genre
type Category struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

// SoftwareCategory is the type for software category
type SoftwareCategory struct {
	ID         int       `json:"-"`
	SoftwareID int       `json:"-"`
	CategoryID int       `json:"-"`
	Category   Category  `json:"category"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
