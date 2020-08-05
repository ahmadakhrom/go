package main

import (
	"chapter-71-exercise-ninja-level-13/test1/chicken"
	"fmt"
)

type animal struct {
	name string
	age  int
}

func main() {

	pelung := animal{
		"Baboon",
		chicken.Years(10),
	}

	fmt.Println(pelung)
	fmt.Println(chicken.YearsTwo(20))

}
