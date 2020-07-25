package main

import "fmt"

func main() {
	var message = make(chan string)		//channels could use to send and receive

	var message = make(<-chan string)    //channle receive only

	//#the way 1
	name := func(s string, d int) {
		k := fmt.Sprintf("Hello %s are you age is %d ?", s, d)
		message <- k
	}

	go name("andi nugraha", 145) //go routine(s) started

	var message = make(chan<- string)    //channle send only
	executed := <-message
	fmt.Println("the way #1 \t", executed)

}