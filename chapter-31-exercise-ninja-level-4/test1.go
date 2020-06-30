package main

import "fmt"

func main() {

	a := [5]int{4, 5, 6, 7, 8}
	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", a)

}
