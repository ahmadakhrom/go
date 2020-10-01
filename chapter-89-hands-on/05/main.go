package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)

	http.Handle("/public/", //setup direct to "/resource" which files have this path
		http.StripPrefix("/public/",
			http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, _ *http.Request){
	err := tpl.ExecuteTemplate(w,"index.gohtml",nil)
	if err != nil {
		http.Error(w,"Error not found",404)
		return
	}
}
