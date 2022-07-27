package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/bartvanbenthem/gofound-blogger/internal/validator"
	"github.com/lib/pq"
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

type PostModel struct {
	DB *sql.DB
}

func (p PostModel) Insert(post *Post) error {
	// Define the SQL query for inserting a new record in the posts table and returning
	// the system-generated data.
	query := `
			INSERT INTO posts (title, content, author, img_urls)
			VALUES ($1, $2, $3, $4)
			RETURNING id, created_at`
	// Create an args slice containing the values for the placeholder parameters from
	// the post struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.
	args := []any{post.Title, post.Content, post.Author, pq.Array(post.ImgURLs)}
	// Use the QueryRow() method to execute the SQL query on our connection pool,
	// passing in the args slice as a variadic parameter and scanning the system-
	// generated id, created_at and version values into the post struct.
	return p.DB.QueryRow(query, args...).Scan(&post.ID, &post.CreatedAt)
}

func (p PostModel) Get(id int64) (*Post, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Define the SQL query for retrieving the post data.
	query := `
		SELECT id, created_at, title, content, author, img_urls
		FROM posts
		WHERE id = $1`
	// Declare a post struct to hold the data returned by the query.
	var post Post
	// Execute the query using the QueryRow() method, passing in the provided id value
	// as a placeholder parameter, and scan the response data into the fields of the
	// post struct. Importantly, notice that we need to convert the scan target for the
	// genres column using the pq.Array() adapter function again.
	err := p.DB.QueryRow(query, id).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.Title,
		&post.Content,
		&post.Author,
		pq.Array(&post.ImgURLs),
	)
	// Handle any errors. If there was no matching post found, Scan() will return
	// a sql.ErrNoRows error. We check for this and return our custom ErrRecordNotFound
	// error instead.
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	// Otherwise, return a pointer to the post struct.
	return &post, nil
}

// Add a placeholder method for updating a specific record in the posts table.
func (p PostModel) Update(post *Post) error {
	return nil
}

// Add a placeholder method for deleting a specific record from the posts table.
func (p PostModel) Delete(id int64) error {
	return nil
}
