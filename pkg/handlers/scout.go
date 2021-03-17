package handlers

import (
	"fmt"
	"net/http"
)

// Scout will do what you think it should do
func Scout() func(w http.ResponseWriter, r *http.Request) {
	// msg := "Reached Scout"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// PICK UP HERE if repo.ConfirmScoutClient

		fmt.Fprintf(w, "This is a test. The request contains: %v", *&r.RequestURI)
	})
}
