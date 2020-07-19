package main

import "fmt"

func main() {

	c := make(chan int)

	foo(c)
	bar(c)
}

func foo(s chan<- int) {
	go func() {
		for i := 1; i <= 100; i++ {
			s <- i
		}
		close(s)
	}()
}

func bar(s chan int) {
	for v := range s {
		fmt.Println(v)
	}
	return
}
