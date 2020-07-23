package main

import "fmt"

func main() {
	var a int
	a++
	fmt.Println(a)
	b := foo()
	fmt.Println(b)

}

func foo() (i int) {
	defer func() { i++ }()
	return 3
}
