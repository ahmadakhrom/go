package main

import "fmt"

func main() {
	a := func(s string) {

		fmt.Println("Hello,", s)
	}

	a("andi nugraha")
	fmt.Printf("%T", a) //type variable "a" is "func (string)"

}
