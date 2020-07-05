package main

import "fmt"

func main() {
	d := foo()
	e := foo()

	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(e())
	fmt.Println(e())

}

func foo() func() int {
	var x int = 100 //------------ I have variable "x" & I was enclosing to block of "func() int"
	return func() int {
		x++
		return x
	}
}
