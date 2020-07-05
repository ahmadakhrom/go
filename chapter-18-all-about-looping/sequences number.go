package main

import "fmt"

func main() {

	h := PrintSequences(5)
	fmt.Println(h)
}

func PrintSequences(n int) int {
	var f int
	for i := 1; i <= n; i++ {
		f, _ = fmt.Println(i)

	}
	return f
}
