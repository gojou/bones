package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	e := run()
	if e != nil {
		log.Fatalf("Fatal error: %v\n", e)
	}
}

func run() (e error) {
	r := mux.NewRouter()
	routes(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	return e
}

// func routes(r *mux.Router) {
// 	r.HandleFunc("/", handlers.Home)
// 	r.HandleFunc("/contact", handlers.Contact)
// 	r.HandleFunc("/about", handlers.About)
// 	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
//
// }
