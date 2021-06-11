package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
	log "github.com/sirupsen/logrus"
)

// CreateAuthor function - POST /authors
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author db.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		utils.Log.Error("Unable to decode request body")
	}

	utils.Log.WithFields(log.Fields{
		"name": author.Name,
	}).Debugf("%s %s - controllers/authors.go - CreateAuthor() -", r.Method, r.URL)

	result, err := db.DB.CreateAuthor(r.Context(), author.Name)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
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

	var listAuthorsParams = db.ListAuthorsParams{
		Limit:  l,
		Offset: o,
	}

	utils.Log.WithFields(log.Fields{
		"limit":  listAuthorsParams.Limit,
		"offset": listAuthorsParams.Offset,
	}).Debugf("%s %s - controllers/authors.go - ListAuthors() -", r.Method, r.URL)

	authors, err := db.DB.ListAuthors(r.Context(), listAuthorsParams)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors)
	}
}

// GetAuthor function - GET /authors/{id}
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	authorID := utils.StrToInt64(id)

	utils.Log.Debugf("%s %s - controllers/authors.go - GetAuthor()", r.Method, r.URL)

	author, err := db.DB.GetAuthor(r.Context(), authorID)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(author)
	}
}

// UpdateAuthor function - PUT /authors/{id}
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	authorID := utils.StrToInt64(id)

	var author db.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		utils.Log.Error("Unable to decode request body")
	}

	utils.Log.WithFields(log.Fields{
		"name": author.Name,
	}).Debugf("%s %s - controllers/authors.go - UpdateAuthor() -", r.Method, r.URL)

	var updateAuthorParams = db.UpdateAuthorParams{
		ID:   authorID,
		Name: author.Name,
	}

	result, err := db.DB.UpdateAuthor(r.Context(), updateAuthorParams)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
