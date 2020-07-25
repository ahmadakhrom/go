package main

import (
	"fmt"
)

func main() {
	d, _ := Sum(1, 2)
	fmt.Println(d)
}

func Sum(i ...int) (int, error) {
	summary := 0

	for _, val := range i {
		summary += val
	}
	return summary, nil
}
