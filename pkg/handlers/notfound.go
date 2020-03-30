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
}
