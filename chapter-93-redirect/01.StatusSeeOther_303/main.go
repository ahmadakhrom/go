package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}
func main() {
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/poo", poo)
	http.ListenAndServe(":8080", nil)
}

func foo(_ http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /: ", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /bar: ", r.Method)
	http.Redirect(w, r, "/", 303) //http.StatusSeeOther
}

func poo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /poo ", r.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}
