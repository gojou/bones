package routing

import (
	"net/http"

	"github.com/gojou/bones/pkg/handlers"
	"github.com/gorilla/mux"
)

// Routes provides a map of the api
func Routes(r *mux.Router) {
	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/scout", handlers.Scout()).Methods("GET")
	r.HandleFunc("/contact", handlers.Contact)
	r.HandleFunc("/contactadd", handlers.ContactAdd)
	r.HandleFunc("/contactlist", handlers.ContactList).Methods("GET")
	r.HandleFunc("/products", handlers.Products).Methods("GET")
	r.HandleFunc("/about", handlers.About).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

}
