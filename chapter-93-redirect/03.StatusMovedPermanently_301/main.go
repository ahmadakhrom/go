package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/favicon.ico",http.NotFoundHandler())
	mux.HandleFunc("/", foo)
	mux.HandleFunc("/bar", bar)
	mux.HandleFunc("/poo", poo)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func foo(_ http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /: ", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /bar: ", r.Method)
	http.Redirect(w, r, "/", 301) //StatusMovedPermanently
}

func poo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /poo ", r.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}
