package main

import "fmt"

func main() {

	f := foo()
	fmt.Println(f())
}

func foo() func() string {
	return func() string {
		d := "hello world"
		return d
	}
}
