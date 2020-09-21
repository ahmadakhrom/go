package main

import (
	"fmt"
	"net/http"
)

type person string

func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("my-key","key from me")
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	w.Header().Add("my-key","hello,world")
	fmt.Fprintln(w,"<h2><em>Hello World, guess who am I ?</em></h2>")
}

func main()  {
var n person
http.ListenAndServe(":8080",n)
}