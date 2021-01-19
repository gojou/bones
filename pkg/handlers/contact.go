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

	if r.Method == "GET" {
		page.Execute(w, nil)
	}

}

//ContactAdd displays the addcontact page
func ContactAdd(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(

		"static/html/_base.html",
		"static/html/contactadd.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
	}

}

//ContactList displays the listcontacts page
func ContactList(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contactlist.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
	}

}
