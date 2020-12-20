package routing

import (
	"net/http"

	"github.com/gojou/bones/pkg/handlers"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/contact", handlers.Contact)
	r.HandleFunc("/contactadd", handlers.ContactAdd)
	r.HandleFunc("/contactlist", handlers.ContactList)
	r.HandleFunc("/products", handlers.Products)
	r.HandleFunc("/about", handlers.About)
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

}
