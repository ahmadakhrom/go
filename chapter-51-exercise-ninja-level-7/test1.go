package main

import "fmt"

func main() {
	a := 120
	fmt.Println("memory address", &a)
	fmt.Println("accessing value with pointer", *&a)
	s := &a
	*s = 110
	fmt.Println("changin value to", *s, "with the same mem address there is", &s)
}

