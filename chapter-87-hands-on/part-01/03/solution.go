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

func index(w http.ResponseWriter, _ *http.Request) {
	data := "hello index"
	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func dog(w http.ResponseWriter, _ *http.Request) {
	data := "hello husky"
	err := tpl.ExecuteTemplate(w, "dog.gohtml", data)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func me(w http.ResponseWriter, _ *http.Request) {
	data := "hello man"
	err := tpl.ExecuteTemplate(w, "dog.gohtml", data)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)   //	mux.HandleFunc("/",index)
	mux.HandleFunc("/dog/", dog) //	mux.HandleFunc("/dog",dog)
	mux.HandleFunc("/me", me)    //	mux.HandleFunc("/me",me)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
