package main

import (
	"io"
	"net/http"
)

type person string

func (p person ) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.URL.Path {
	case "/man":
		io.WriteString(w,"<h1>hello i'm a man</h1>")
	case "/girl":
		io.WriteString(w,"hello i'm a girl")
	}
}

func main()  {
	var n person
	http.ListenAndServe(":8080",n)
}