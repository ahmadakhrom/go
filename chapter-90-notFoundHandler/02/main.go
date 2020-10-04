package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())ls

	http.ListenAndServe(":8080", nil)

	
}

func foo(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Println(r.URL.Path)
	fmt.Fprintln(w, "look at terminal")
}
