package main

import (
	// Provides access to SQL database operations
	"database/sql"
	// Helps encode Go structs to JSON and decode JSON to Go structs (used for API requests/responses)
	"encoding/json"
	// print logs
	"fmt"
	// Standard logging utility (used for logging errors or info in terminal)
	"log"
	// Provides HTTP server and request/response utilities (core of API handling)
	"net/http"
	//to convert str to int or vice versa
	"strconv"
	// This is the MSSQL driver for database/sql — connects Go to SQL Server (must import anonymously)
	_ "github.com/denisenkom/go-mssqldb"
	// route library
	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"` // json  is it handles encodeing and decoding request and response
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *sql.DB

// connecting to db
func initDB() {
	//connString := "sqlserver://sa:YOUR_PASSWORD@localhost:1433?database=BookDB"
	connString := "sqlserver://localhost:1433?database=BookDB&trusted_connection=yes"
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Ping failed:", err)
	}
	fmt.Println("Connected to SQL Server")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	stmt := `INSERT INTO Books (Title, Author) OUTPUT INSERTED.ID VALUES (@p1, @p2)`
	err := db.QueryRow(stmt, book.Title, book.Author).Scan(&book.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT ID, Title, Author FROM Books")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		rows.Scan(&b.ID, &b.Title, &b.Author)
		books = append(books, b)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var book Book
	err := db.QueryRow("SELECT ID, Title, Author FROM Books WHERE ID = @p1", id).
		Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		http.Error(w, "Book not found", 404)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	_, err := db.Exec("UPDATE Books SET Title = @p1, Author = @p2 WHERE ID = @p3", book.Title, book.Author, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	book.ID, _ = strconv.Atoi(id)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := db.Exec("DELETE FROM Books WHERE ID = @p1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Book with ID %s deleted", id)
}

func main() {
	initDB()
	router := mux.NewRouter()

	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
