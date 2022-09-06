package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Register the relevant methods, URL patterns and handler functions for our
	// healthcheck endpoint using the HandlerFunc() method.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// posts endpoints using the HandlerFunc() method.
	router.HandlerFunc(http.MethodGet, "/v1/posts", app.listPostsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/posts", app.createPostHandler)
	router.HandlerFunc(http.MethodGet, "/v1/posts/:id", app.showPostHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/posts/:id", app.updatePostHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/posts/:id", app.deletePostHandler)
	// Return the httprouter instance.
	return router
}
