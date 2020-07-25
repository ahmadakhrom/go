package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := 10
	b := 10
	c := 100

foo(a, b,c)


}

func foo(numGoroutines, numData, TotalLoop int) {
	d := make(chan int)

	//go func(){
	for i := 1; i <= numGoroutines; i++ {
		go func() {
			for i := 1; i <= numData; i++ {
				d <- i
			}
		}()
		fmt.Println("num of gor-----", runtime.NumGoroutine())
	}
	//}()
	for i := 1; i <= TotalLoop; i++ {
		fmt.Println(i, <-d)
	}
	close(d)
	return
}
