package main

import (
	"html/template"
	"log"
	"net/http"
)

type person int

var tpl *template.Template

func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("%s",err)
	}
	tpl.ExecuteTemplate(w,"index.gohtml",r.Form)
}

func init()  {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main()  {
	var s person
	http.ListenAndServe(":8080",s)
}