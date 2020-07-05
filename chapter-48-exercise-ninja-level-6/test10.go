package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)
	a++
	fmt.Println(a)

	{ //----------------------------enclosed
		var b int = 12               //1
		fmt.Println("var b is :", b) //12
	}

	var b int
	fmt.Println(b) //0

	g := foo()

	fmt.Println(g()) //51

}

func foo() func() int {
	return func() int {
		c := 50
		c++
		return c
	}
}
