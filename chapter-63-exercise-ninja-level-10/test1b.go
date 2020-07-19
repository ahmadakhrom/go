package main

import "fmt"

func main() {
	c := make(chan int, 1) //"1" channel is buffered

	go func() {
		c <- 42
		c <- 43
	}()
	fmt.Println(<-c) //executed
	fmt.Println(<-c) //executed
}
