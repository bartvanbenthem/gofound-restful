package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bartvanbenthem/gofound-blogger/internal/data"
)

func (app *application) createPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new blog post")
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
