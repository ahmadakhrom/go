package main

import "fmt"

func main() {
	d := foo()
	fmt.Printf("%T \n", d)

	e := d()
	fmt.Println(e)
	f := foo()() //the way for returning "func()" in "func foo"
	fmt.Println(f)
}

func foo() func() int {
	return func() int {
		return 876
	}
}
