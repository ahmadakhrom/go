package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/",foo)

	fs := http.FileServer(http.Dir("./public"))
	log.Fatal(http.ListenAndServe(":8080", fs))
}

func foo(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "", nil)
	if err != nil {
		http.Error(w, "Error Not Foundss", 404)
		return
	}
}