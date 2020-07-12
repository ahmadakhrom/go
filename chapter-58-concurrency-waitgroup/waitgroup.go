package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("os\t\t", runtime.GOOS)
	fmt.Println("architecture\t", runtime.GOARCH)
	fmt.Println("numb cpu\t", runtime.NumCPU())
	fmt.Println("num goroutine\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("numb cpu\t", runtime.NumCPU())
	fmt.Println("num goroutine\t", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func bar() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
