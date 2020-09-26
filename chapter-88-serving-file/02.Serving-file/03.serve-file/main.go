package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/sunset.jpg", bar)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "/sunset.jpg">`)
}

func bar(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../source/sunset.jpg")
}
