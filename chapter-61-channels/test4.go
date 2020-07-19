package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)


close(q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}

func receive(a, b <-chan int) {

	for {
		select {
		case res := <-a:
			fmt.Println("-----", res)

		case i, err := <-b:
			if err == true {
				fmt.Println("-----", i, err)
				return
			} else {
				fmt.Println("-----", i, err)
				return
			}
		}

	}
}
