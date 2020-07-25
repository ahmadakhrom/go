package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	b := gen(a)

	receive(b, a)
	fmt.Println("about to exit")
}

func gen(a chan<- int) <-chan int {
	c := make(chan int)
	var i int
	go func() {
		for i = 0; i < 100; i++ {
			c <- i
		}
		//a<-1  //other way below to stopping looping
		a <- i //for stopping looping until reached max number "100"
		close(c)
	}()
	return c
}

func receive(a, b <-chan int) {

	for {
		select {
		case res := <-a:
			fmt.Println("-----", res)
		case <-b:
			return
		}
	}
}
