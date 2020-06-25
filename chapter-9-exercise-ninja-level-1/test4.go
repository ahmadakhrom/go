package main

import "fmt"

type digit int

var x digit

func main() {

	fmt.Println(x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Println(x)
}
