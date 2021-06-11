package main

import (
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/routes"
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

	router := mux.NewRouter()
	apiSubrouter := router.PathPrefix("/api").Subrouter()
	routes.InitializeBookRoutes(apiSubrouter)
	routes.InitializeAuthorRoutes(apiSubrouter)
	utils.Log.Info("Finished initializing routes")

	utils.Log.Info("Server ready and listening on port " + utils.Config.Port)
	utils.Log.Fatal(http.ListenAndServe(":"+utils.Config.Port, router))
}
