package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
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
	router := httprouter.New()
	router.Handler("GET","/favicon.ico",http.NotFoundHandler())
	router.GET("/", foo)
	router.POST("/bar", bar)
	router.GET("/poo", poo)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func foo(_ http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /: ", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /bar: ", r.Method)
	http.Redirect(w, r, "/", 301) //StatusMovedPermanently
}

func poo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println(time.Now().Format("15:04:05"),"request method /poo ", r.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}
