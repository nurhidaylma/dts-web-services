package endpoints

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Desc   string
}

var books = []Book{
	{ID: 1, Title: "How Are You?", Author: "Bimal Jalal", Desc: "Book 1"},
	{ID: 2, Title: "A Place Called Home", Author: "Preety Shenoy", Desc: "Book 2"},
	{ID: 3, Title: "Queen of Fire", Author: "Devika Rangachari", Desc: "Book 3"},
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(books)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func getBookByID(w http.ResponseWriter, r *http.Request)  {
	
}
