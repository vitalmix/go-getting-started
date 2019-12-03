package main
import (
	"net/http"
	"log"
	"html/template"
	_ "github.com/lib/pq"
	"database/sql"
	"os"
	"fmt"
)
var db *sql.DB
func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}
		books, err := dbGetBooks()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, books)
	}
}
func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		name := r.Form.Get("name")
		year := r.Form.Get("year")
		length := r.Form.Get("length")
		err := dbAddBook(name, year, length)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println(port)
	}
	return ":" + port
}
func main() {
	err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addBookHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}
