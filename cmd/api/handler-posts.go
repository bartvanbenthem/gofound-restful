package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bartvanbenthem/gofound-restful/internal/data"
	"github.com/bartvanbenthem/gofound-restful/internal/validator"
)

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Author  string   `json:"author,omitempty"`
		ImgURLs []string `json:"img_urls,omitempty"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &data.Post{
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
		ImgURLs: input.ImgURLs,
	}

	// Initialize a new Validator.
	v := validator.New()

	// Call the Validatepost() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidatePost(v, post); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Call the Insert() method on our post model, passing in a pointer to the
	// validated post struct. This will create a record in the database and update the
	// post struct with the system-generated information.
	err = app.models.Posts.Insert(post)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// When sending a HTTP response, we want to include a Location header to let the
	// client know which URL they can find the newly-created resource at. We make an
	// empty http.Header map and then use the Set() method to add a new Location header,
	// interpolating the system-generated ID for our new post in the URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posts/%d", post.ID))
	// Write a JSON response with a 201 Created status code, the post data in the
	// response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"post": post}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showPostHandler(w http.ResponseWriter, r *http.Request) {
	// read param from URL
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// DB query
	post, err := app.models.Posts.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"post": post}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the post ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Fetch the existing post record from the database, sending a 404 Not Found
	// response to the client if we couldn't find a matching record.
	post, err := app.models.Posts.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Declare an input struct to hold the expected data from the client.
	var input struct {
		Title   *string  `json:"title"`
		Content *string  `json:"content"`
		Author  *string  `json:"author,omitempty"`
		ImgURLs []string `json:"img_urls,omitempty"`
	}
	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the request body to the appropriate fields of the post record
	if input.Title != nil {
		post.Title = *input.Title
	}

	if input.Content != nil {
		post.Content = *input.Content
	}
	if input.Author != nil {
		post.Author = *input.Author
	}
	if input.ImgURLs != nil {
		post.ImgURLs = input.ImgURLs // we don't need to dereference a slice.
	}

	// Validate the updated post record, sending the client a 422 Unprocessable Entity
	// response if any checks fail.
	v := validator.New()
	if data.ValidatePost(v, post); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Pass the updated post record to our new Update() method.
	err = app.models.Posts.Update(post)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Write the updated post record in a JSON response.
	err = app.writeJSON(w, http.StatusOK, envelope{"post": post}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the post ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Delete the post from the database, sending a 404 Not Found response to the
	// client if there isn't a matching record.
	err = app.models.Posts.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Return a 200 OK status code along with a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "post successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPostsHandler(w http.ResponseWriter, r *http.Request) {
	// To keep things consistent with our other handlers, we'll define an input struct
	// to hold the expected values from the request query string.
	var input struct {
		Title   string
		ImgURLs []string
		data.Filters
	}

	// Initialize a new Validator instance.
	v := validator.New()
	// Call r.URL.Query() to get the url.Values map containing the query string data.
	qs := r.URL.Query()
	// Use our helpers to extract the title and genres query string values, falling back
	// to defaults of an empty string and an empty slice respectively if they are not
	// provided by the client.
	input.Title = app.readString(qs, "title", "")
	input.ImgURLs = app.readCSV(qs, "img_urls", []string{})
	// Read the page and page_size query string values into the embedded struct.
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	// Read the sort query string value into the embedded struct.
	input.Filters.Sort = app.readString(qs, "sort", "id")
	// Add the supported sort values for this endpoint to the sort safelist.
	input.Filters.SortSafelist = []string{"id", "title", "author", "-id", "-title", "-year", "-author"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Call the GetAll() method to retrieve the movies, passing in the various filter
	// parameters.
	posts, err := app.models.Posts.GetAll(input.Title, input.ImgURLs, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"posts": posts}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
