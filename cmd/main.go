package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	e := run()
	if e != nil {
		log.Fatalf("Fatal error: %v\n", e)
	}
}

func run() error {
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":8888", r)
	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to the Index!</h1>")
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Welcome to the Contact!</h1>")
	}

}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Not all who wander are lost. But you are. 404</h1>")

}
