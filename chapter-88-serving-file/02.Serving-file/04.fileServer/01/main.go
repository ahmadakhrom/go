package main

import (
	"io"
	"net/http"
)

func main() {
	// diretory on (".") is root
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.HandleFunc("/foo/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	//from dir root func foo could be accessing dir /assets/image
	io.WriteString(w,`<img src="/assets/sunset.jpg">`)
}

