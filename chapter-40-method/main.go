package main

import "fmt"

type weapon struct {
	model string
}

type terorist struct {
	name string
	age  int
	weapon
}

func (w weapon) SayHello() {
	fmt.Println("model weapon these terorist is", w.model)
}

func (t terorist) Target() {
	fmt.Println("he named was", t.name, "and ages is", t.age)
}

func main() {
	//show values of field
	target := terorist{
		weapon: weapon{
			model: "AK47",
		},
		name: "Osama Bin Laden",
		age:  56,
	}
	fmt.Println(target)

	target.SayHello()
	target.Target()

}
