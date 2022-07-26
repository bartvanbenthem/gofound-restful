package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bartvanbenthem/gofound-blogger/internal/data"
	"github.com/bartvanbenthem/gofound-blogger/internal/validator"
)

func (app *application) createPostsHandler(w http.ResponseWriter, r *http.Request) {
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

	post := &data.Post{Title: input.Title,
		Content: input.Content,
		Author:  input.Author,
		ImgURLs: input.ImgURLs}

	// Initialize a new Validator.
	v := validator.New()

	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidatePost(v, post); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showPostsHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	post := data.Post{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "test-post",
		Author:    "",
		Content:   "this is a test post",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"post": post}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
