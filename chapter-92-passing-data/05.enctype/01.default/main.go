package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

type person struct { //use when your applying form value method from request
	First  string
	Last   string
	Active bool
}

func main() {
	http.HandleFunc("/", foo)
	//http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	//use when your applying form value method from request
	//first := r.FormValue("First")
	//last := r.FormValue("Last")
	//active := r.FormValue("Active") == "true"
	//data := person{first, last, active,}

	byteSlice := make([]byte, r.ContentLength)
	r.Body.Read(byteSlice)
	bodyContent := string(byteSlice)

	err := tpl.ExecuteTemplate(w, "index.html", bodyContent)
	//err := tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), 500) //status internal server error
	}
}
