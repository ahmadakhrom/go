package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/applyProcess", applyProcess)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8080", nil)
}

func errorHttp(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "error occured", 500) //internal server error
		log.Fatal(err)                      //or os.exit(2)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
		errorHttp(w, err)
	}
}c

func contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
		errorHttp(w, err)
	}
}

func apply(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
		errorHttp(w, err)
	}
}

func applyProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "apply-process.gohtml", nil)
		errorHttp(w, err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
		errorHttp(w, err)
	}
}
