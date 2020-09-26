package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/",foo)
	http.HandleFunc("/",foo)
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, _ *http.Request)  {
	w.Header().Set("Content-Type","text/html; charset=utf-8")

	io.WriteString(w, `
		<!--image doesn't serve'-->
		<img src="../source/sunset.jpg">
	`)
}
