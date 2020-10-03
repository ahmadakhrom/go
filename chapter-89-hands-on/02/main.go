package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	http.HandleFunc("/", foo)

	fs := http.FileServer(http.Dir("."))
	log.Fatal(http.ListenAndServe(":8080", fs))
}

func foo(w http.ResponseWriter, _ *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "error 404", 404)
		return
	}
}
