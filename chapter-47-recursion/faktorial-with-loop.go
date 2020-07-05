package main

import "fmt"

func main() {
	//fmt.Println(5 * 4 * 3 * 2 * 1)

	d := faktorialLoop(5)
	fmt.Println(d)
}

func faktorialLoop(n int) int {

	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
