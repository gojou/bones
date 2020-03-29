package handlers

import (
	"fmt"
	"net/http"
)

// NotFound handles 404
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Not all who wander are lost. But you are. 404</h1>")

}
