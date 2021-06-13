package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
	log "github.com/sirupsen/logrus"
)

type createAuthorRequest struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type updateAuthorRequest struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

// CreateAuthor function - POST /authors
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var req createAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Log.WithFields(log.Fields{
		"name": req.Name,
	}).Debugf("%s %s - controllers/authors.go - CreateAuthor() -", r.Method, r.URL)

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

	result, err := db.DB.CreateAuthor(r.Context(), req.Name)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// ListAuthors function - GET /authos
func ListAuthors(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}
	l := utils.StrToInt32(limit)

	offset := r.URL.Query().Get("offset")
	if offset == "" {
		offset = "0"
	}
	o := utils.StrToInt32(offset) * l

	utils.Log.WithFields(log.Fields{
		"limit":  l,
		"offset": o,
	}).Debugf("%s %s - controllers/authors.go - ListAuthors() -", r.Method, r.URL)

	authors, err := db.DB.ListAuthors(r.Context(), db.ListAuthorsParams{
		Limit:  l,
		Offset: o,
	})
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}

// GetAuthor function - GET /authors/{id}
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	authorID := utils.StrToInt64(id)

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}
	l := utils.StrToInt32(limit)

	offset := r.URL.Query().Get("offset")
	if offset == "" {
		offset = "0"
	}
	o := utils.StrToInt32(offset) * l

	utils.Log.WithFields(log.Fields{
		"limit":  l,
		"offset": o,
	}).Debugf("%s %s - controllers/authors.go - GetAuthor() -", r.Method, r.URL)

	author, err := db.DB.GetAuthor(r.Context(), db.GetAuthorParams{
		ID:     authorID,
		Limit:  l,
		Offset: o,
	})
	if author.ID == 0 {
		utils.Log.Errorf("AuthorID: %s does not exist!", id)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

// UpdateAuthor function - PUT /authors/{id}
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	authorID := utils.StrToInt64(id)

	var req updateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Log.WithFields(log.Fields{
		"name": req.Name,
	}).Debugf("%s %s - controllers/authors.go - UpdateAuthor() -", r.Method, r.URL)

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

	result, err := db.DB.UpdateAuthor(r.Context(), db.UpdateAuthorParams{
		ID:   authorID,
		Name: req.Name,
	})
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// DeleteAuthor function - DELETE /authors/{id}
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	authorID := utils.StrToInt64(id)

	utils.Log.Debugf("%s %s - controllers/authors.go - DeleteAuthor()", r.Method, r.URL)

	if err := db.DB.DeleteAuthor(r.Context(), authorID); err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
