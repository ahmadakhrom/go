package main

import "fmt"

func main() {

	//loop in loop
	lenQuestion := 5
	lensubQuestion := 3
	for i := 1; i <= lenQuestion; i++ {
		fmt.Println("Question ", i)
		for i := 1; i <= lensubQuestion; i++ {
			fmt.Println("\tcontains have a sub question as many is ", i)
		}
	}

	//make a statement "if + condition" inside "for{}"
	a := 1
	for {
		if a > 3 {
			break
		}
		fmt.Println(a)
		a++
	}
	fmt.Println("looping is done..")
}

//
