package main

import "fmt"

func main() {
	ch := make(chan interface{})

	go func() {
		ch <- 12.2
		ch <- "hello world"
		close(ch)
	}()

	resp1, err := <-ch
	resp2, err := <-ch
	resp3, err := <-ch

	fmt.Println(resp1, err)
	fmt.Println(resp2, err)
	fmt.Println(resp3, err)
}
