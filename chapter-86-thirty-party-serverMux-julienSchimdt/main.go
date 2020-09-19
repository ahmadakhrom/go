package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:year/:month/:article-name", blogRead)
	mux.POST("/blog/:year/:month/:article-name", blogWrite) // not work cause method is post or write request

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Error occured : %s", err)
	}

}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), 500) //http.StatusInternalServerError
		log.Fatalf("%s", err)
	}
}

func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

//
func applyProcess(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply-process.gohtml", nil)
	HandleError(w, err)
}

func user(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, " User %s \n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Read Year %s \n", ps.ByName("year"))
	fmt.Fprintf(w, "Read Month %s \n", ps.ByName("month"))
	fmt.Fprintf(w, "Read article-name %s \n", ps.ByName("article-name"))
}

func blogWrite(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Write Year %s \n", ps.ByName("category"))
	fmt.Fprintf(w, "Write Month %s \n", ps.ByName("article"))
	fmt.Fprintf(w, "Write article-name %s \n", ps.ByName("article-name"))
}
