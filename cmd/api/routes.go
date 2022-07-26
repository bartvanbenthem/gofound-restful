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
	// endpoints using the HandlerFunc() method.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/posts", app.createPostsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/posts/:id", app.showPostsHandler)
	// Return the httprouter instance.
	return router
}
