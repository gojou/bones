package main

import (
	"log"
	"net/http"

	"github.com/gojou/bones/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	e := run()
	if e != nil {
		log.Fatalf("Fatal error: %v\n", e)
	}
}

func run() (e error) {

	port := ":8080"
	r := mux.NewRouter()

	routes(r)
	log.Printf("Starting web server on port %v\n", port)
	http.ListenAndServe(port, r)
	return e
}

func routes(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/contact", handlers.Contact)
	r.HandleFunc("/about", handlers.About)
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

}
