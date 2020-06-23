package main

import "fmt"

var x, y, z = 42, "james bond", true

func main() {

	fmt.Println(x)                       //a single print statements
	fmt.Println(y)                       //a single print statements
	fmt.Println(z)                       //a single print statements
	fmt.Println("print all : ", x, y, z) //multiple print statements

	fmt.Printf("type of x is : %T \n", x)
	fmt.Printf("type of y is : %T \n", y)
	fmt.Printf("type of z is : %T \n", z)

}

//you can acces the doc these materi on my course
//on https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
