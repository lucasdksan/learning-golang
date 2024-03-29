package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {
	templates = template.Must(templates.ParseGlob("*.html"))
	u := user{"Lucas da Silva", "lokasmega@gmail.com"}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", u)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
