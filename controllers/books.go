package controllers

import (
	"encoding/json"
	"net/http"

	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
	log "github.com/sirupsen/logrus"
)

// CreateBook function - POST /books
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book db.Book
	json.NewDecoder(r.Body).Decode(&book)

	utils.Log.WithFields(log.Fields{
		"title": book.Title,
	}).Debugf("%s %s - controllers/books.go - CreateBook() -", r.Method, r.URL)

	result, err := db.DB.CreateBook(r.Context(), book.Title)
	if err != nil {
		utils.Log.Error(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
