package main

import "fmt"

func main() {
	var i interface{} = 12

	fmt.Printf("%T", i) //check TYPE of var

	s := i.(int) //you cannot declare var s to TYPE of int
	fmt.Println(s)

	s = i.(float64) //you cannot declare var s to TYPE of float64
	fmt.Println(s)

}
