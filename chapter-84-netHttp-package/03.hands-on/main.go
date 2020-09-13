package main

import (
	"html/template"
	"log"
	"net/http"
)

type greeting string

var tpl *template.Template

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("%s", err)
	}
	tpl. ExecuteTemplate(w,"index.gohtml",r.Form)
}

func init()  {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var t greeting
	http.ListenAndServe(":8080",t)
}
