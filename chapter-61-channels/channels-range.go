package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go foo(ch, 15)
	bar(ch)

	//===========================================================
	// ch = make(chan int)
	// go func() {
	// 	for j := 1; j <= 10; j++ {
	// 		ch <- j
	// 	}
	// 	close(ch)
	// }()

	// for val := range ch {
	// 	fmt.Println(val)

	// }
	// //===========================================================
	// fmt.Println("\nwell done!")
}

func foo(ch chan int, num int) {
	for i := 1; i <= num; i++ {
		ch <- i
	}
	close(ch)
}

func bar(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
