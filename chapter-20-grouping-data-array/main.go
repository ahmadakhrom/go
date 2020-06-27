package main

import "fmt"

func main() {

	var a [10]int
	fmt.Println("length of var a is :", a)

	a[9] = 1
	fmt.Println(a)
}
