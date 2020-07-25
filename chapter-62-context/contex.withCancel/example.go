package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//print err ctx & num of gouroutine used
	fmt.Println("error check 1 : ", ctx.Err())
	fmt.Println("num of goroutines 1 : ", runtime.NumGoroutine())

	go func() {
		n := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200) //every 1/5 second do print loping
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2) //sleep after 2 second running
	fmt.Println("error check 2 : ", ctx.Err())
	fmt.Println("num of goroutines 2 : ", runtime.NumGoroutine())

	fmt.Println("ready to cancel")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2) //sleep after 2 second running
	fmt.Println("error check 3 : ", ctx.Err())
	fmt.Println("num of goroutines 3 : ", runtime.NumGoroutine())
}
