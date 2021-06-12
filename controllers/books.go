package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
	log "github.com/sirupsen/logrus"
)

// CreateBook function - POST /books
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var createBookParams db.CreateBookParams
	if err := json.NewDecoder(r.Body).Decode(&createBookParams); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.Log.WithFields(log.Fields{
		"title":     createBookParams.Title,
		"author_id": createBookParams.AuthorID,
	}).Debugf("%s %s - controllers/books.go - CreateBook() -", r.Method, r.URL)

	book, err := db.DB.CreateBook(r.Context(), createBookParams)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	}
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
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}

// GetBook function - GET /books/{id}
func GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookID := utils.StrToInt64(id)

	utils.Log.Debugf("%s %s - controllers/books.go - GetBook()", r.Method, r.URL)

	book, err := db.DB.GetBook(r.Context(), bookID)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	}
}

// UpdateBook function - PUT /books/{id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookID := utils.StrToInt64(id)

	var book db.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.Log.Error("Unable to decode request body")
	}

	utils.Log.WithFields(log.Fields{
		"title": book.Title,
	}).Debugf("%s %s - controllers/books.go - UpdateBook() -", r.Method, r.URL)

	var updateBookParams = db.UpdateBookParams{
		ID:    bookID,
		Title: book.Title,
	}

	result, err := db.DB.UpdateBook(r.Context(), updateBookParams)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

// DeleteBook function - DELETE /books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookID := utils.StrToInt64(id)

	utils.Log.Debugf("%s %s - controllers/books.go - DeleteBook()", r.Method, r.URL)

	if err := db.DB.DeleteBook(r.Context(), bookID); err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
