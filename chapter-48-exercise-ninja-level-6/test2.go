package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6}

	d := foo(numbers...)
	fmt.Println(d)

	e := bar(numbers)
	fmt.Println(e)
}

func foo(n ...int) int {
	i := 0
	for _, v := range n {
		i += v
	}
	return i
}

func bar(n []int) int {
	i := 0 //-------------------------------??????????????????????
	for _, v := range n {
		i += v
	}
	return i
}
