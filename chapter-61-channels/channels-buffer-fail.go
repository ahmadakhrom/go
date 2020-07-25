package main

import "fmt"

func main() {

	var ch = make(chan int, 2) //-----second parameter of var ch, must be filled more than 1 to pass values that channel

	ch <- 911
	ch <- 108

	fmt.Println("---- \t", <-ch)
	fmt.Println("---- \t", <-ch)
	fmt.Printf("---- \t%T", ch)
}
