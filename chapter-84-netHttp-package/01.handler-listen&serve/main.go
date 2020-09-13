package main

import (
	"fmt"
	"net/http"
)

type animal int	//make struct

func (d animal) ServeHTTP(w http.ResponseWriter, r *http.Request) {  //make method from X struct
	fmt.Fprintln(w,"your cat so cute")
}

func main()  {
	var z animal
	http.ListenAndServe(":8080",z)
}