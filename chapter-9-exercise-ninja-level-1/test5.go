package main

import "fmt"

type digit int

var x int
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T \n", x)
}
