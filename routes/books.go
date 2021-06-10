package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moeabdol/bookstore-api-golang/controllers"
)

// BooksRouter function
func BooksRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books", controllers.ListBooks).Methods(http.MethodGet)
	return r
}
