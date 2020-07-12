package main

import (
	"fmt"
	"runtime"
	"sync"
	_ "time"
)

func main() {

	fmt.Println("num CPUs\t:", runtime.NumCPU())
	fmt.Println("num goroutine\t:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	counter := 0
	const maxL = 10000
	wg.Add(maxL)

	for i := 1; i <= maxL; i++ {
		go func() {

			v := counter
			runtime.Gosched()
			v++
			counter = v

			wg.Done()
		}()
		fmt.Println("num goroutine---", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("num CPUs\t:", runtime.NumCPU())
	fmt.Println("num goroutine\t:", runtime.NumGoroutine())
	fmt.Println("num counter\t:", counter)
}
