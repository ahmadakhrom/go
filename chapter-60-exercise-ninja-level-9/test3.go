package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("num goroutines :", runtime.NumGoroutine())

	var wg sync.WaitGroup

	const count = 100
	var counter = 0
	wg.Add(count)

	for i := 1; i <= count; i++ {
		go func() {

			v := counter

			runtime.Gosched()
			v++
			counter = v

			wg.Done()
		}()
		fmt.Println("--queue goroutines ", i, "--num goroutines ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("num goroutines :", runtime.NumGoroutine())
	fmt.Println("count :", counter)
}
