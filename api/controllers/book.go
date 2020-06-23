package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ono5/book-list/api/utils"

	"github.com/ono5/book-list/api/models"
)

type Controller struct{}

var books []models.Book

// GetBooks - get all books
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		rows, err := db.Query("select * from books")
		utils.LogFatal(err)

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			utils.LogFatal(err)

			books = append(books, book)
		}

		json.NewEncoder(w).Encode(books)
	}
}
