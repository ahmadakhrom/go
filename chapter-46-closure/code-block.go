package main

import "fmt"

func main() {
	var a int = 12
	fmt.Println(a)
	a++

	//code block of variable b
	{
		var b int = 14
		fmt.Println(b)
	}

	fmt.Println(b) // undefined b, coz b not define in ouside of scope b / in scope main func
}
