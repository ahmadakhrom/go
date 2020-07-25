package main

import "fmt"

func main() {

	var message = make(chan string)

	//-------------#the way 1
	name := func(s string, d int) {
		k := fmt.Sprintf("Hello %s are you age is %d ?", s, d)
		message <- k
	}

	go name("andi nugraha", 145) //go routine(s) started

	executed := <-message
	fmt.Println("the way #1 \t", executed)

	//-------------#the way 2
	var ch = make(chan int)
	go func() {
		ch <- 911
	}()

	fmt.Println("the way #2 \t", <-ch)

}
