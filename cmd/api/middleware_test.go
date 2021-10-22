package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestEnableCORS(t *testing.T) {
	var myH myHandler

	h := enableCORS(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestTokenVerify(t *testing.T) {
	var myH myHandler

	h := TokenVerify(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
