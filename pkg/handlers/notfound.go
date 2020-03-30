package handlers

import (
	"html/template"
	"net/http"
)

// NotFound handles 404
func NotFound(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/404.html",
	))

	if r.Method == "GET" {
		w.WriteHeader(http.StatusNotFound)

		page.Execute(w, nil)
		//	return
	}

	// w.Header().Set("Content-Type", "text/html")
	// w.WriteHeader(http.StatusNotFound)
	// fmt.Fprint(w, "<h1>Not all who wander are lost. But you are. 404</h1>")

}
