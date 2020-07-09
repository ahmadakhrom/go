package main

import "fmt"

func main() {

	a := "hello world"
	fmt.Println("value from memory addres of a :", a)
	fmt.Println("memory addres of a :", &a)
	foo(&a)

}

func foo(s *string) {
	*s = "hello mars"
	fmt.Println("changing value a...")
	fmt.Println("memory addres of a :", *s)
	fmt.Println("value from memory addres of a :", s)
}
