package main

import (
	"chapter-69-exercise-ninja-level-12/godark"
	//"github.com/ahmadakhrom/go/chapter-69-exercise-ninja-level-12/godark"
	"fmt"

)

type Person struct {
	name string
	age int
}


func main() {
	
	data := Person{
		"John Wick",
		godark.Years(10),
	}

	fmt.Println(data)
		fmt.Println(data.name)
			fmt.Println(data.age)


}