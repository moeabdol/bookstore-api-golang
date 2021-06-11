package main

import (
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/routes"
	"github.com/moeabdol/bookstore-api-golang/utils"
	"github.com/spf13/viper"
)

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

	router := mux.NewRouter()
	apiSubrouter := router.PathPrefix("/api").Subrouter()
	routes.InitializeBookRoutes(apiSubrouter)
	routes.InitializeAuthorRoutes(apiSubrouter)

	port := viper.Get("PORT").(string)
	utils.Log.Info("Server ready and listening on port " + port)
	utils.Log.Fatal(http.ListenAndServe(":"+port, router))
}
