package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (a person) speak() {
	fmt.Println("hallo,", a.first, a.last+".", "you are ages,", a.age)
}

func main() {

	user := person{
		first: "andi",
		last:  "nugraha",
		age:   27,
	}

	user.speak()
}
