package main

import (
	"fmt"
)

func main() {

	//one variable the way for use append
	a := []int{1, 2, 3, 4}
	a = append(a, 5, 6, 7, 8)
	fmt.Println(a)

	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9}

	//	fmt.Println(append(x, y)) <== error
	// if you append with difference variable you must be add a dot three times "..." after y
	// and must be same TYPE
	fmt.Println(append(x, y...))

}
