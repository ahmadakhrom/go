package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/sunset.jpg", bar)

	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "/sunset.jpg">`)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fileSource, err := os.Open("../source/sunset.jpg")
	defer fileSource.Close()

	if err != nil {
		http.Error(w, "Error 404", http.StatusNotFound)
		return
	}

	fileInfo, err := fileSource.Stat()
	if err != nil {
		http.Error(w, "error 404", http.StatusNotFound)
		return
	}

	http.ServeContent(w, r, fileSource.Name(), fileInfo.ModTime(), fileSource)
}
