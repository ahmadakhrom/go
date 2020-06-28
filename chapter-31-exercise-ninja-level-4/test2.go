package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}

	for i, v := range a {
		fmt.Println(" on index", i, "and value is", v)
	}

	fmt.Printf("and TYPE of data is %T", a)
}
