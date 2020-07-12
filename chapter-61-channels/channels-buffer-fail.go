package main

import "fmt"

func main() {

	var ch = make(chan int, 1) //-----second parameter of var ch, must be filled more than 1 to pass values that channel

	ch <- 911
	ch <- 108

	fmt.Println("the way #2 \t", <-ch)
	fmt.Println("the way #2 \t", <-ch)
}
