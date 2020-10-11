package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func checkErr500(w http.ResponseWriter, e error)  {
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
}

func checkErr404(w http.ResponseWriter, e error)  {
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
}

func foo(w http.ResponseWriter, r *http.Request) {

	var data string
	if r.Method == http.MethodPost {
		//open requested file
		multipartFile, fileHeader, err := r.FormFile("q")
		checkErr404(w, err)
		defer multipartFile.Close()

		//read file
		bytesData, err := ioutil.ReadAll(multipartFile)
		checkErr500(w, err)
		data = string(bytesData)

		//store file in the server
		directoryString, err := os.Create(filepath.Join("./user/", fileHeader.Filename))
		checkErr500(w, err)
		defer directoryString.Close()

		//build / write file to folder
		_, err = directoryString.Write(bytesData)
		checkErr500(w, err)
	}
	//if used a template(S) no need for set header below:
	w.Header().Set("Content-Type","text/html;charset=utf-8")

	tpl.ExecuteTemplate(w, "index.gohtml",data)
}

//package main
//
//import (
//	"fmt"
//	"html/template"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"path/filepath"
//)
//
//var tpl *template.Template
//
//func init() {
//	tpl = template.Must(template.ParseGlob("templates/*"))
//}
//
//func main() {
//	http.HandleFunc("/", foo)
//	http.Handle("/favicon.ico", http.NotFoundHandler())
//	http.ListenAndServe(":8080", nil)
//}
//
//func foo(w http.ResponseWriter, req *http.Request) {
//
//	var s string
//	if req.Method == http.MethodPost {
//
//		// open
//		f, h, err := req.FormFile("q")
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		defer f.Close()
//
//		// for your information
//		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
//
//		// read
//		bs, err := ioutil.ReadAll(f)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		s = string(bs)
//
//		// store on server
//		dst, err := os.Create(filepath.Join("./user/", h.Filename))
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		defer dst.Close()
//
//		_, err = dst.Write(bs)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
//
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	tpl.ExecuteTemplate(w, "index.gohtml", s)
//}