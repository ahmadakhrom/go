package main

import "fmt"

func main() {

	fmt.Println(5 * 4 * 3 * 2 * 1) //120

	n := faktorial(5)
	fmt.Println(n)
}

func faktorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
	//return n * faktorial(5-1) * faktorial(4-1) * faktorial(3-1) * faktorial(2-1) * 1

}
