package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() chan int {
	c := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(ch chan int) {

	for i := range ch {
		fmt.Println("-----", i)
	}
}
