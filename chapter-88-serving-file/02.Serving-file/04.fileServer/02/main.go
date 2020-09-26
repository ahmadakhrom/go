package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/resource/", http.StripPrefix("/resource",http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `<img src = "/resource/sunset.jpg">`)
}
