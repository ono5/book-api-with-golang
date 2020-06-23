package controllers

import (
	"database/sql"
	"net/http"

	"github.com/ono5/book-list/api/models"
	"github.com/ono5/book-list/api/repository/bookrepository"
	"github.com/ono5/book-list/api/utils"
)

type Controller struct{}

var books []models.Book

// GetBooks - get all books
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		books = []models.Book{}
		bookRepo := bookrepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)
		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}
