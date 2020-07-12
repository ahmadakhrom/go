package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var max = 100

func main() {
	fmt.Println("num of Goroutines started\t", runtime.NumGoroutine())
	fmt.Println("num of CPUs started\t\t", runtime.NumCPU())
	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("num of Goroutines mid\t", runtime.NumGoroutine())
	fmt.Println("num of CPUs mid\t\t", runtime.NumCPU())

	wg.Wait()

	fmt.Println("num of Goroutines end\t", runtime.NumGoroutine())
	fmt.Println("num of CPUs end\t\t", runtime.NumCPU())

}

func foo() {
	for i := 1; i <= max; i++ {
		fmt.Println("goroutines 1", i)
	}
	fmt.Println("hello wordl 1")
	wg.Done()
}

func bar() {
	for i := 1; i <= max; i++ {
		fmt.Println("goroutines 2", i)
	}
	fmt.Println("hello wordl 2")
	wg.Done()
}
