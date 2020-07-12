package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) speak() {
	fmt.Println("hello I am", (*p).Name, "and", (*p).Age, "years old")
}

type human interface {
	speak()
}

func SaySomething(h human) {
	h.speak()
}

func main() {

	p1 := Person{
		"andi nugraha",
		27,
	}

	SaySomething(&p1)
	//SaySomething(p1) //-----------cannot cal without "&" coz saysomething, it was pointer type
	fmt.Println(p1.Name, p1.Age)
}
