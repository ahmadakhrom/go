package main

import (
	"context"
	"fmt"
	"runtime"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel wher all fihished

	fmt.Println("num gouroutines", runtime.NumGoroutine())

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 3 {
			break
		}
	}

	fmt.Println("num gouroutines", runtime.NumGoroutine())
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)

	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // retur to avoid leak memory by goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
