package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main()  {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	var a string
	fmt.Println(r.Method)

	if r.Method == http.MethodPost {
		//open file
		file, fileHeader, err := r.FormFile("q")
		{
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			defer file.Close()
		}
		//information
		fmt.Println("file :", file, "\nfile header :", fileHeader, "\nerr :", err)
		//read file
		byteFile, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		a = string(byteFile)
	}
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>` +a)
}