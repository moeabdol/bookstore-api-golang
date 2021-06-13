package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
	log "github.com/sirupsen/logrus"
)

type createUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

// CreateUser function - POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Log.WithFields(log.Fields{
		"username": req.Username,
		"password": req.Password,
		"email":    req.Email,
	}).Debugf("%s %s - controllers/users.go - CreateUser() -", r.Method, r.URL)

	valErrors := utils.ValidateStruct(req)
	if len(valErrors) != 0 {
		utils.Log.Errorf("Validation errors: %s", valErrors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string][]string{
			"errors": valErrors,
		})
		return
	}

	count, err := db.DB.EmailExists(r.Context(), req.Email)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if count > 0 {
		utils.Log.Errorf("Email: %s already exists!", req.Email)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"errors": "User already exists!",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := db.DB.CreateUser(r.Context(), db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	})
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser function - GET /users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := utils.StrToInt64(id)

	utils.Log.Debugf("%s %s - controllers/users.go - GetUser()", r.Method, r.URL)

	user, err := db.DB.GetUser(r.Context(), userID)
	if user.ID == 0 {
		utils.Log.Errorf("User with id: %s does not exist!", id)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
