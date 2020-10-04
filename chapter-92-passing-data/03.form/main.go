package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

type person struct {
	FirstName string
	LastName string
	Status bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request)  {
	a := r.FormValue("firstName")
	b := r.FormValue("lastName")
	c := r.FormValue("status") == "on"

	err := tpl.ExecuteTemplate(w, "index.html",person{
		FirstName: a,
		LastName:  b,
		Status:    c,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
