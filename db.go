package lab5
import (
	"database/sql"
	"fmt"
	_ "os"
)
type Book struct{
	Name string
	Year string
	Length string
}
const (
	DB_USER = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME = "test"
)
func dbConnect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS books (book_name text,book_year text,book_length text)"); err != nil {
		return err
	}
	return nil
}
func dbAddBook(name, year, length string) error {
	sqlstmt := "INSERT INTO books VALUES ($1, $2, $3)"
	_, err := db.Exec(sqlstmt, name, year, length)
	if err != nil {
		return err
	}
	return nil
}
func dbGetBooks() ([]Book, error) {
	var books []Book
	stmt, err := db.Prepare("SELECT book_name, book_year, book_length FROM books")
	if err != nil {
		return books, err
	}
	res, err := stmt.Query()
	if err != nil {
		return books, err
	}
	var tempBook Book
	for res.Next() {
		err = res.Scan(&tempBook.Name, &tempBook.Year, &tempBook.Length)
		if err != nil {
			return books, err
		}
		books = append(books, tempBook)
	}
	return books, err
}
