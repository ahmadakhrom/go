package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main() {
	insideMax := 10
	outsideMax := 10

x := foo(insideMax,outsideMax)

for i := 1; i <= 100; i++ {
	fmt.Println(i,<-x)
}
	

}

func foo(insideMax, outsideMax int) chan int{
a := make(chan int)
var wg sync.WaitGroup

wg.Add(1)
	for i := 1; i <= 10; i++ { //num of goroutines
		go func()  {
			for i := 1; i <= 100; i++ {
				a <- i
			}
		}()
		fmt.Println("num of gor",runtime.NumGoroutine())
	}
	return a
}