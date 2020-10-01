package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("bar.gohtml"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar/", bar)

	http.Handle("/resource/", //setup direct to "/resource" which files have this path
		http.StripPrefix("/resource",
			http.FileServer(http.Dir("./assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}

func foo(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html:charset=utf-8")
	w.Header().Set("foo", "ran")
}

func bar(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "bar.gohtml", nil)
	w.Header().Set("response", "template executed")
}
