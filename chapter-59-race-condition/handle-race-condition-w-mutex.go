package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("num CPUs", runtime.NumCPU())
	fmt.Println("num goroutine", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0
	const maxL = 100
	wg.Add(maxL)

	for i := 1; i <= maxL; i++ {
		go func() {
			mu.Lock()

			v := counter
			runtime.Gosched()
			v++
			counter = v

			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("num CPUs", runtime.NumCPU())
	fmt.Println("num goroutine", runtime.NumGoroutine())
	fmt.Println("num counter", counter)
}
