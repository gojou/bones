package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gojou/bones/cmd/routing"
	"github.com/gorilla/mux"
)

func main() {

	e := start()
	if e != nil {
		log.Fatalf("Fatal error: %v\n", e)
	}
}

func start() (e error) {
	r := mux.NewRouter()
	routing.Routes(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	return e
}
