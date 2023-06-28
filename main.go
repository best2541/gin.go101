package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title string `title`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write(([]byte("<h1>Hello</h1>")))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	book := Book{
		Title: "test",
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/book", GetBook)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
