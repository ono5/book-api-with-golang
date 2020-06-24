package main

import (
	"log"
	"net/http"

	"database/sql"
	_ "database/sql"

	_ "github.com/lib/pq"
	"github.com/ono5/book-list/api/controllers"
	"github.com/ono5/book-list/api/driver"
	"github.com/ono5/book-list/api/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var books []models.Book
var db *sql.DB

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requeste-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
