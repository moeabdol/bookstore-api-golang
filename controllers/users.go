package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
	log "github.com/sirupsen/logrus"
)

// CreateUser function - POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var args db.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.Log.WithFields(log.Fields{
		"username": args.Username,
		"password": args.Password,
		"email":    args.Email,
	}).Debugf("%s %s - controllers/users.go - CreateUser() -", r.Method, r.URL)

	user, err := db.DB.CreateUser(r.Context(), args)
	if err != nil {
		utils.Log.Error(err)
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code.Name() {
			case "unique_violation":
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "User already exists!",
				})
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// GetUser function - GET /users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := utils.StrToInt64(id)

	utils.Log.Debugf("%s %s - controllers/users.go - GetUser()", r.Method, r.URL)

	user, err := db.DB.GetUser(r.Context(), userID)
	if user.Email == "" {
		w.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
