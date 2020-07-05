package main

import "fmt"

func main() {

	var y int

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
