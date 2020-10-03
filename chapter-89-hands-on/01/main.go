package main

import (
	"html/template"
	"io"
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
	http.HandleFunc("/sunset.jpg", doo)


	//http.Handle("/resource/", //setup direct to "/resource" which files have this path
	//	http.StripPrefix("/resource",
	//		http.FileServer(http.Dir("./assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}

func foo(w http.ResponseWriter, _ *http.Request) {
	//w.Header().Set("Content-Type", "text/html:charset=utf-8")
	//w.Header().Set("foo", "ran")
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "bar.gohtml", nil)
	if err != nil {
		http.Error(w, "error occured code 500", 500)
	}
	//w.Header().Set("response", "template executed")
}

func doo(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w, r,"./assets/sunset.jpg")
}