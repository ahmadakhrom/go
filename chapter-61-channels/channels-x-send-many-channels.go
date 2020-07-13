package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("starting..")

	e := make(chan int)
	o := make(chan int)
	fanIn := make(chan int)
	num := 500

	//sending channels
	go Send(e, o, num)
	go Receive(e, o, fanIn)

	for val := range fanIn {
		fmt.Println(val)
	}

	fmt.Println("done!")

}

//sending channels
func Send(e, o chan<- int, num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

//receive channels
func Receive(e, o <-chan int, fanIn chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanIn <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)

}
