package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/bartvanbenthem/gofound-restful/internal/validator"
	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author,omitempty"`
	ImgURLs   []string  `json:"img_urls,omitempty"`
	Version   int64     `json:"version"`
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
			RETURNING id, created_at, version`
	// Create an args slice containing the values for the placeholder parameters from
	// the post struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.
	args := []any{post.Title, post.Content, post.Author, pq.Array(post.ImgURLs)}
	// Use the QueryRow() method to execute the SQL query on our connection pool,
	// passing in the args slice as a variadic parameter and scanning the system-
	// generated id, created_at and version values into the post struct.

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return p.DB.QueryRowContext(ctx, query, args...).Scan(&post.ID, &post.CreatedAt, &post.Version)
}

func (p PostModel) Get(id int64) (*Post, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Define the SQL query for retrieving the post data.
	query := `
		SELECT id, created_at, title, content, author, img_urls, version
		FROM posts
		WHERE id = $1`
	// Declare a post struct to hold the data returned by the query.
	var post Post

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Make sure to cancel the context before the method returns.
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.Title,
		&post.Content,
		&post.Author,
		pq.Array(&post.ImgURLs),
		&post.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

// Add a placeholder method for updating a specific record in the posts table.
func (p PostModel) Update(post *Post) error {
	// Declare the SQL query for updating the record and returning the new version
	// number.
	query := `
			UPDATE posts
			SET title = $1, content = $2, author = $3, img_urls = $4, version = version + 1
			WHERE id = $5 AND version = $6
			RETURNING version`
	// Create an args slice containing the values for the placeholder parameters.
	args := []any{
		post.Title,
		post.Content,
		post.Author,
		pq.Array(post.ImgURLs),
		post.ID,
		post.Version,
	}

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Execute the SQL query. If no matching row could be found, we know the movie
	// version has changed (or the record has been deleted) and we return our custom
	// ErrEditConflict error.
	err := p.DB.QueryRowContext(ctx, query, args...).Scan(&post.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

// Add a placeholder method for deleting a specific record from the posts table.
func (p PostModel) Delete(id int64) error {
	// Return an ErrRecordNotFound error if the post ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}
	// Construct the SQL query to delete the record.
	query := `
			DELETE FROM posts
			WHERE id = $1`

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Execute the SQL query using the Exec() method, passing in the id variable as
	// the value for the placeholder parameter. The Exec() method returns a sql.Result
	// object.
	result, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	// Call the RowsAffected() method on the sql.Result object to get the number of rows
	// affected by the query.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// If no rows were affected, we know that the posts table didn't contain a record
	// with the provided ID at the moment we tried to delete it. In that case we
	// return an ErrRecordNotFound error.
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (p PostModel) GetAll(title string, imageUrls []string, filters Filters) ([]*Post, Metadata, error) {
	// Construct the SQL query to retrieve all records.
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, created_at, title, content, author, img_urls, version
		FROM posts
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (img_urls @> $2 OR $2 = '{}')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// As our SQL query now has quite a few placeholder parameters, let's collect the
	// values for the placeholders in a slice. Notice here how we call the limit() and
	// offset() methods on the Filters struct to get the appropriate values for the
	// LIMIT and OFFSET clauses.
	args := []any{title, pq.Array(imageUrls), filters.limit(), filters.offset()}

	// Use QueryContext() to execute the query. This returns a sql.Rows resultset
	// containing the result.
	rows, err := p.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	posts := []*Post{}

	for rows.Next() {

		var post Post

		err := rows.Scan(
			&totalRecords,
			&post.ID,
			&post.CreatedAt,
			&post.Title,
			&post.Content,
			&post.Author,
			pq.Array(&post.ImgURLs),
			&post.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	// Generate a Metadata struct, passing in the total record count and pagination
	// parameters from the client.
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return posts, metadata, nil
}
