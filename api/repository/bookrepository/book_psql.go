package bookrepository

import (
	"database/sql"

	"github.com/ono5/book-list/api/models"
)

type BookRepository struct{}

// GetBooks - get all books from db
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("select * from books")
	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}
	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}
