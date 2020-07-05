package main

import "fmt"

var y int

func main() {

	fmt.Println(y)
	y++
	fmt.Println(y)
	d := bar(y)
	fmt.Println(d)
}

func bar(x int) int {
	x++
	return x
}
