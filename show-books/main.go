package main

import (
	"html/template"
	"net/http"
)

// var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var tmpls = template.Must(template.ParseFiles("templates/index.html"))

type Book struct {
	Title  string
	Author string
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	if err := tmpls.ExecuteTemplate(w, "index.html", book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
