package main

import (
	"io"
	"net/http"
	"os"
)

func main()  {
	http.HandleFunc("/",foo)
	http.HandleFunc("/sunset.jpg",bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="/sunset.jpg">`)  //request serving a file to dir /sunset.jpg
}

func bar(w http.ResponseWriter, r *http.Request)  {  //serving requested file
	fileSource, err := os.Open("../source/sunset.jpg")

	if err != nil {
		http.Error(w,"eror code 404",http.StatusNotFound)
		return
	}
	defer fileSource.Stat()

	io.Copy(w, fileSource)
}