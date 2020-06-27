package main

import "fmt"

const (
	a = 1
	b = 3.3
	c = "john doe"
)

func main() {
	//a = 9 //error cannot assign '9' to variable const 'a'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T %T %T", a, b, c)

}
