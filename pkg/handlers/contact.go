package handlers

import (
	"fmt"
	"net/http"
)

//Contact displays the contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Welcome to the Contact!</h1>")
	}
}
