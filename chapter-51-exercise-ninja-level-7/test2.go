package main

import "fmt"

type person struct {
	name string
	age  int
}

func (a *person) changeMe() {
	fmt.Println("your name is", (*a).name, "and", *&a.age, "years old")
}

//-----------------------------------------------------------------------------------

func changeYou(d *person) (string, int) {
	(*d).name = "angga setiawan"
	(*d).age = 28
	s := (*d).name
	f := (*d).age
	return s, f
}

func main() {

	person := person{
		name: "andi nugraha",
		age:  27,
	}
	fmt.Println(person)
	person.changeMe()

	//-----------------------------------------------------------------------------------

	d, e := changeYou(&person)
	fmt.Println(d, e)
}
