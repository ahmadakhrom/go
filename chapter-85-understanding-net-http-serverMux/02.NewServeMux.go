package main

import (
	"io"
	"net/http"
)

type person int
type animal int
type dafault int

func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello husky duggy lucky..")
}

func (p animal) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello olla holla lala..")
}

func (p dafault) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "this dafault..")
}

func main()  {

	var a person
	var b animal
	var c dafault


	mux := http.NewServeMux()
	mux.Handle("/person/",a)
	mux.Handle("/animal",b)
	mux.Handle("/",c)

http.ListenAndServe(":8080",mux)
}