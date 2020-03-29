package handlers

import (
	"fmt"
	"net/http"
)

// Home displays the default "/" page
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to the Index!</h1>")
	}
}
