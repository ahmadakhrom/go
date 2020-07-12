package main

import "fmt"

func main() {

	c := make(chan string)

	var say string
	fmt.Println("input below")
	fmt.Scanln(&say)

	go foo(say, c) //send channel

	fmt.Println("\nchannel printed :")

	bar(c) //receive channel

}

func foo(s string, c chan<- string) {
	c <- s
}

func bar(c <-chan string) {
	fmt.Println(<-c)
}
