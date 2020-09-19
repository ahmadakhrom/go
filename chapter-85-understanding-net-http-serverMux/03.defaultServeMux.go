package main

import (
	"io"
	"net/http"
)

// ------------------------ bug example ------------------------
//type person int
//func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "hello person")
//}
//
//type animal int
//func (a animal) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "hello animal")
//}
//
//type dfault int
//func (d dfault) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "hello world")
//}

//func main()  {
//	var p person
//	var  a animal
//	var d dfault
//
//	http.Handle("/",d)
//	http.Handle("/person",p)
//	http.Handle("/animal",a)
//
//	http.ListenAndServe(":8080",nil)
//}

func person(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello person")
}

func animal(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello animal")
}

func dfault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello dfault handle")
}

func main() {
	http.HandleFunc("/person", person)
	http.HandleFunc("/animal", animal)
	http.HandleFunc("/a", dfault)

	//when you want to use just http.Handle, you can convert be like :
	//http.Handle("/person",http.HandleFunc(person))
	//http.Handle("/animal",http.HandleFunc(animal))
	//http.Handle("/dfault",http.HandleFunc(dfault))

	http.ListenAndServe(":8080", nil)

}
