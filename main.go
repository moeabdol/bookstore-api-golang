package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/routes"
	"github.com/moeabdol/bookstore-api-golang/utils"
	"github.com/spf13/viper"
)

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}

func main() {
	utils.ReadConfig()
	utils.InitializeLogger()
	utils.Log.Info("Finished reading .env config file")

	err := db.ConnectToDatabase()
	if err != nil {
		utils.Log.Fatalf("Not able to connect to database %s", err)
	} else {
		utils.Log.Info("Connected to database")
	}

	r := mux.NewRouter()
	mount(r, "/api", routes.BooksRouter())

	port := viper.Get("PORT").(string)
	utils.Log.Info("Server ready and listening on port " + port)
	utils.Log.Fatal(http.ListenAndServe(":"+port, r))
}
