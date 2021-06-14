package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
	log "github.com/sirupsen/logrus"
)

type signupRequest struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

type signinRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Signup function - POST /users
func Signup(w http.ResponseWriter, r *http.Request) {
	var req signupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Log.WithFields(log.Fields{
		"username": req.Username,
		"password": strings.Repeat("*", len(req.Password)),
		"email":    req.Email,
	}).Debugf("controllers/auth.go - Signup() -")

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

// Signin function - POST /auth/signin
func Signin(w http.ResponseWriter, r *http.Request) {
	var req signinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Log.WithFields(log.Fields{
		"email":    req.Email,
		"password": strings.Repeat("*", len(req.Password)),
	}).Debugf("controllers/auth.go - Signin() -")

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
	} else if count == 0 {
		utils.Log.Errorf("Email: %s does not exists!", req.Email)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"errors": "Email does not exist!",
		})
		return
	}

	user, err := db.DB.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = utils.ComparePassword(req.Password, user.Password)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.CreateToken(user.Username, time.Minute*5, "keys/private-key.pem")
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": accessToken,
	})
}
