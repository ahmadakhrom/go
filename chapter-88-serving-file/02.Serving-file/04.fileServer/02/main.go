package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/resource/",  //setup direct to "/resource" which files have this path
		http.StripPrefix("/resource",
			http.FileServer(http.Dir("./assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}


func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	//which path is begin prefix with "/resource" will be direct to directory assets
	io.WriteString(w, `<img src = "/resource/sunset.jpg">`)
}
