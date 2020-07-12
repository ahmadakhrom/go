package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("num CPUs", runtime.NumCPU())
	fmt.Println("num goroutine", runtime.NumGoroutine())

	var counter int64

	const maxL = 100
	var wg sync.WaitGroup
	wg.Add(maxL)

	for i := 1; i <= maxL; i++ {
		go func() {

			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter\t", atomic.LoadInt64(&counter))

			wg.Done()
		}()
		fmt.Println("num goroutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("num goroutine", runtime.NumGoroutine())
	fmt.Println("num counter", counter)
}
