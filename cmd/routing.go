package main

import (
	"net/http"

	"github.com/gojou/bones/pkg/handlers"
	"github.com/gorilla/mux"
)

func routing(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/about", handlers.About)
	r.HandleFunc("/contact", handlers.Contact).Methods("GET")
	r.HandleFunc("/contact/list", handlers.ContactList).Methods("GET")
	r.HandleFunc("/contact/add", handlers.ContactAdd).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

}
