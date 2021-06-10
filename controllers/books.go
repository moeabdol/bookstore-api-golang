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
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.Log.Error("Unable to decode request body")
	}

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

// ListBooks function - GET /books
func ListBooks(w http.ResponseWriter, r *http.Request) {
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

	var listBooksParams = db.ListBooksParams{
		Limit:  l,
		Offset: o,
	}

	utils.Log.WithFields(log.Fields{
		"limit":  listBooksParams.Limit,
		"offset": listBooksParams.Offset,
	}).Debugf("%s %s - controllers/books.go - ListBooks() -", r.Method, r.URL)

	books, err := db.DB.ListBooks(r.Context(), listBooksParams)
	if err != nil {
		utils.Log.Error(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
