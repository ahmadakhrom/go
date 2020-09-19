package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	data := "hello index"
	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	data := "hello husky"
	err := tpl.ExecuteTemplate(w, "dog.gohtml", data)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	data := "hello man"
	err := tpl.ExecuteTemplate(w, "dog.gohtml", data)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me", me)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
