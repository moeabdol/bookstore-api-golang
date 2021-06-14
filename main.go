package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moeabdol/bookstore-api-golang/controllers"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/middleware"
	"github.com/moeabdol/bookstore-api-golang/utils"
)

func main() {
	utils.LoadConfig()
	utils.InitializeLogger()

	db.ConnectToDatabase()
	utils.Log.Info("Connected to database " + utils.Config.DBName)

	r := mux.NewRouter()
	r.Use(middleware.RequestLogger)

	// User routes
	authRoutes := r.PathPrefix("/api/auth").Subrouter()
	authRoutes.HandleFunc("/signup", controllers.Signup).Methods(http.MethodPost)
	authRoutes.HandleFunc("/signin", controllers.Signin).Methods(http.MethodPost)

	userRoutes := r.PathPrefix("/api/users").Subrouter()
	userRoutes.HandleFunc("", controllers.CreateUser).Methods(http.MethodPost)
	userRoutes.HandleFunc("/{id}", controllers.GetUser).Methods(http.MethodGet)

	// Book routes
	bookRoutes := r.PathPrefix("/api/books").Subrouter()
	bookRoutes.HandleFunc("", controllers.CreateBook).Methods(http.MethodPost)
	bookRoutes.HandleFunc("", controllers.ListBooks).Methods(http.MethodGet)
	bookRoutes.HandleFunc("/{id}", controllers.GetBook).Methods(http.MethodGet)
	bookRoutes.HandleFunc("/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	bookRoutes.HandleFunc("/{id}", controllers.DeleteBook).Methods(http.MethodDelete)

	// Author routes
	authorRoutes := r.PathPrefix("/api/authors").Subrouter()
	authorRoutes.HandleFunc("", controllers.CreateAuthor).Methods(http.MethodPost)
	authorRoutes.HandleFunc("", controllers.ListAuthors).Methods(http.MethodGet)
	authorRoutes.HandleFunc("/{id}", controllers.GetAuthor).Methods(http.MethodGet)
	authorRoutes.HandleFunc("/{id}", controllers.UpdateAuthor).Methods(http.MethodPut)
	authorRoutes.HandleFunc("/{id}", controllers.DeleteAuthor).Methods(http.MethodDelete)
	utils.Log.Info("Finished initializing routes")

	utils.Log.Info("Server ready and listening on port " + utils.Config.Port)
	utils.Log.Fatal(http.ListenAndServe(":"+utils.Config.Port, r))
}
