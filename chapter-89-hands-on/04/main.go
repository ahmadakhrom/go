package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {

	http.HandleFunc("/", foo)

	http.Handle("/resources/", //setup direct to "/resource" which files have this path
		http.StripPrefix("/resources/",
			http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, _ *http.Request) {
	err := tpl.Execute(w,nil)
	if err != nil {
		http.Error(w, "error 500", 500)
		return
	}
}
