package handlers

import (
	"html/template"
	"net/http"
)

//Contact displays the contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact.html",
	))

	page.Execute(w, nil)

}

//ContactAdd displays the "add a contact" page
func ContactAdd(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact_add.html",
	))

	page.Execute(w, nil)

}

//ContactList displays the "List Contacts" page
func ContactList(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact_list.html",
	))

	page.Execute(w, nil)

}
