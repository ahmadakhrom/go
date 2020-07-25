package main

import "fmt"

func main() {
	n, err := fmt.Println("234")
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	fmt.Println(n)
	fmt.Println("all be allright")
}
