package main

import (
	"encoding/json"
	"log"
	"net/http"

	"database/sql"
	_ "database/sql"

	_ "github.com/lib/pq"
	"github.com/ono5/book-list/api/controllers"
	"github.com/ono5/book-list/api/driver"
	"github.com/ono5/book-list/api/models"
	"github.com/ono5/book-list/api/utils"

	"github.com/gorilla/mux"
)

var books []models.Book
var db *sql.DB

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	utils.LogFatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)

	utils.LogFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	utils.LogFatal(err)

	rowsUpdated, err := result.RowsAffected()
	utils.LogFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id = $1", params["id"])
	utils.LogFatal(err)

	rowsDeleted, err := result.RowsAffected()
	utils.LogFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}
