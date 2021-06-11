package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moeabdol/bookstore-api-golang/controllers"
)

// InitializeBookRoutes function
func InitializeBookRoutes(r *mux.Router) {
	r.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books", controllers.ListBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
}
