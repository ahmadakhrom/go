package main

import (
	"fmt"
	"go/types"
)

//format : keyword identifier type
type gangster struct {
	name string
	age  int
}

type status struct {
	gangster
	active bool
}

func (w gangster) speak() {
	fmt.Println("we have soldier there is :", w.name, w.age)
}

func (s status) speak() {
	fmt.Println("we have soldier there is----", s.name, s.age)
}

type attacker interface {
	speak()
}

func foo(a attacker) {

	switch a.(type) {
	case gangster:
		fmt.Println("hallo..", a.(gangster).name)

	case status:
		fmt.Println("hallo..xx", a.(status).name)
	}
	fmt.Println("passed aa & bb", a)
}

//type(keyword) anothertype(identifier) int(type)
type anotherType int

func main() {
	//==================================================
	aa := status{
		gangster: gangster{
			"andy nugraha",
			30,
		},
		active: true,
	}

	bb := status{
		gangster: gangster{
			"budi budiman",
			38,
		},
		active: true,
	}

	cc := gangster{
		name: "alex muhadi",
		age:  40,
	}

	fmt.Println("1", aa) //{{andy 30} true}
	fmt.Println("2", bb) //{{budi budiman 38} true}
	fmt.Println("3", cc) //{budi 40}

	aa.speak() //we have soldier there is---- andy 30
	bb.speak() //we have soldier there is---- budi budiman 38

	foo(aa) //hallo..xx andy //passed aa & bb {{andy 30} true}
	foo(bb) //hallo..xx budi budiman //passed aa & bb {{budi budiman 38} true}
	foo(cc) //hallo.. budi //passed aa & bb {budi 40}

	var x anotherType = 102
	fmt.Println("value from var x with type of anthorType is :", x)
	fmt.Printf("type of variable x is %T \n", x)

	y := int(x)
	fmt.Println(y)
	fmt.Printf("type of variable y is %T \n", y)
}
