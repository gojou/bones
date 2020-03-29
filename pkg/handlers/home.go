package handlers

import (
	"html/template"
	"net/http"
)

// Home displays the default "/" page
func Home(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/index.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		//	return
	}

}
