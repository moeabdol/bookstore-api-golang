package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moeabdol/bookstore-api-golang/controllers"
	"github.com/moeabdol/bookstore-api-golang/utils"
)

// InitializeAuthorRoutes function
func InitializeAuthorRoutes(r *mux.Router) {
	r.HandleFunc("/authors", controllers.CreateAuthor).Methods(http.MethodPost)

	utils.Log.Info("Finished initilizing author routes")
}
