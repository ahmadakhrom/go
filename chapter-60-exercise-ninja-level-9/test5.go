package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("num goroutines :", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var counter int64

	const count = 100

	wg.Add(count)

	for i := 1; i <= count; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()

			wg.Done()
		}()

		fmt.Println("--queue goroutines:", i, "--num goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("num goroutines :", runtime.NumGoroutine())
	fmt.Println("count :", counter)
}
