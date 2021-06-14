package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moeabdol/bookstore-api-golang/controllers"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
)

func main() {
	utils.LoadConfig()
	utils.InitializeLogger()

	db.ConnectToDatabase()
	utils.Log.Info("Connected to database " + utils.Config.DBName)

	r := mux.NewRouter()
	// User routes
	r.HandleFunc("/api/users", controllers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/users/{id}", controllers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/api/users/login", controllers.LoginUser).Methods(http.MethodPost)

	// Book routes
	r.HandleFunc("/api/books", controllers.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/api/books", controllers.ListBooks).Methods(http.MethodGet)
	r.HandleFunc("/api/books/{id}", controllers.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods(http.MethodDelete)

	// Author routes
	r.HandleFunc("/api/authors", controllers.CreateAuthor).Methods(http.MethodPost)
	r.HandleFunc("/api/authors", controllers.ListAuthors).Methods(http.MethodGet)
	r.HandleFunc("/api/authors/{id}", controllers.GetAuthor).Methods(http.MethodGet)
	r.HandleFunc("/api/authors/{id}", controllers.UpdateAuthor).Methods(http.MethodPut)
	r.HandleFunc("/api/authors/{id}", controllers.DeleteAuthor).Methods(http.MethodDelete)
	utils.Log.Info("Finished initializing routes")

	utils.Log.Info("Server ready and listening on port " + utils.Config.Port)
	utils.Log.Fatal(http.ListenAndServe(":"+utils.Config.Port, r))
}
