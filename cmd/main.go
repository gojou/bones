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

	r.HandleFunc("/", handler)
	http.ListenAndServe(":8888", r)
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to the jungle!</h1>")
	}

}
