package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, r *http.Request)  {
	println(r.URL.Path)
	fmt.Fprintln(w, "look at terminal")
}
