package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moeabdol/bookstore-api-golang/controllers"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
)

func main() {
	utils.ReadConfig()
	utils.InitializeLogger()
	utils.Log.Info("Finished reading .env config file")

	err := db.ConnectToDatabase()
	if err != nil {
		utils.Log.Fatalf("Not able to connect to database %s", err)
	} else {
		utils.Log.Info("Connected to database " + utils.Config.DbName)
	}

	r := mux.NewRouter()
	r.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books", controllers.ListBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/authors", controllers.CreateAuthor).Methods(http.MethodPost)
	r.HandleFunc("/authors", controllers.ListAuthors).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id}", controllers.GetAuthor).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id}", controllers.UpdateAuthor).Methods(http.MethodPut)
	r.HandleFunc("/authors/{id}", controllers.DeleteAuthor).Methods(http.MethodDelete)
	utils.Log.Info("Finished initializing routes")

	utils.Log.Info("Server ready and listening on port " + utils.Config.Port)
	utils.Log.Fatal(http.ListenAndServe(":"+utils.Config.Port, r))
}
