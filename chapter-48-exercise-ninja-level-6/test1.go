package main

import "fmt"

func main() {

	f := foo()
	i,s  := bar()

	fmt.Println(f)
	fmt.Println(i, s)

}

func foo() int {
	n := 12
	return n
}

func bar() (int, string) {
	a := 18
	b := "andy nugraha"
	return a, b
}
