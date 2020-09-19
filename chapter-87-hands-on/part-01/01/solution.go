package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "hello index")
}

func dog(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "hello husky")
}

func me(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "hello man")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me", me)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
