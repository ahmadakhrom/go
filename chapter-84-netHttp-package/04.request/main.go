package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type fruit string

type urlRequest struct {
	Method string
	URL *url.URL
	Submissions map[string][]string
	Header http.Header
	Host string
	ContentLength int64
}

func (f fruit) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("%s",err)
	}

	data := urlRequest{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.gohtml",data)
}

func init()  {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main()  {
	var s fruit
	http.ListenAndServe(":8080",s)
}