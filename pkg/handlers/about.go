package handlers

import (
	"fmt"
	"net/http"
)

//About displays the contact page
func About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/about" {
		fmt.Fprint(w, "<h1>Welcome to the About!</h1>")
	}
}
