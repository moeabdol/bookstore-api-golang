package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/moeabdol/bookstore-api-golang/db/sqlc"
	"github.com/moeabdol/bookstore-api-golang/utils"
	log "github.com/sirupsen/logrus"
)

type createBookRequest struct {
	Title    string `json:"title" validate:"required,min=3,max=255"`
	AuthorID int64  `json:"author_id" validate:"required"`
}

type updateBookRequest struct {
	Title    string `json:"title" validate:"required,min=3,max=255"`
	AuthorID int64  `json:"author_id" validate:"required"`
}

// CreateBook function - POST /books
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var req createBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Log.WithFields(log.Fields{
		"title":     req.Title,
		"author_id": req.AuthorID,
	}).Debugf("%s %s - controllers/books.go - CreateBook() -", r.Method, r.URL)

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

	book, err := db.DB.CreateBook(r.Context(), db.CreateBookParams{
		Title:    req.Title,
		AuthorID: req.AuthorID,
	})
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
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

	utils.Log.WithFields(log.Fields{
		"limit":  l,
		"offset": o,
	}).Debugf("%s %s - controllers/books.go - ListBooks() -", r.Method, r.URL)

	books, err := db.DB.ListBooks(r.Context(), db.ListBooksParams{
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
	json.NewEncoder(w).Encode(books)
}

// GetBook function - GET /books/{id}
func GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookID := utils.StrToInt64(id)

	utils.Log.Debugf("%s %s - controllers/books.go - GetBook()", r.Method, r.URL)

	book, err := db.DB.GetBook(r.Context(), bookID)
	if book.ID == 0 {
		utils.Log.Errorf("Book with id: %s does not exist!", id)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook function - PUT /books/{id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookID := utils.StrToInt64(id)

	var req updateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Log.Error("Unable to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Log.WithFields(log.Fields{
		"title":     req.Title,
		"author_id": req.AuthorID,
	}).Debugf("%s %s - controllers/books.go - UpdateBook() -", r.Method, r.URL)

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

	count, err := db.DB.BookTitleExists(r.Context(), req.Title)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if count > 0 {
		utils.Log.Errorf("Book: %s already exists!", req.Title)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"errors": "Book already exists!",
		})
		return
	}

	count, err = db.DB.AuthorIDExists(r.Context(), req.AuthorID)
	if err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if count == 0 {
		utils.Log.Errorf("AuthorID: %s does not exist!", fmt.Sprint(req.AuthorID))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"errors": "Author does not exist!",
		})
		return
	}

	result, err := db.DB.UpdateBook(r.Context(), db.UpdateBookParams{
		ID:       bookID,
		Title:    req.Title,
		AuthorID: req.AuthorID,
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

// DeleteBook function - DELETE /books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookID := utils.StrToInt64(id)

	utils.Log.Debugf("%s %s - controllers/books.go - DeleteBook()", r.Method, r.URL)

	if err := db.DB.DeleteBook(r.Context(), bookID); err != nil {
		utils.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
