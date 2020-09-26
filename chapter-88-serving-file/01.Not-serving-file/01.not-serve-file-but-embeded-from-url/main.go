package main

import (
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/",foo)
	http.ListenAndServe(":8080",nil)
	
}

func foo(w http.ResponseWriter, _ *http.Request)  {
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	w.Header().Set("Another-Key","lkujsufcio8eyt4i87yvn")

	io.WriteString(w,` 
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}


