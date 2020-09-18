package main

import (
	"io"
	"net/http"
)

type person int
func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello person")
}

type animal int
func (a animal) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello animal")
}

type dfault int
func (d dfault) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func main()  {
	var p person
	var  a animal
	var d dfault

	http.Handle("/",d)
	http.Handle("/person",p)
	http.Handle("/animal",a)

	http.ListenAndServe(":8080",nil)
}