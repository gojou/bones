package handlers

import (
	"html/template"
	"net/http"
)

//Products displays the products page
func Products(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/products.html",
	))

	page.Execute(w, nil)

}
