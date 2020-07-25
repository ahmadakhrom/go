package main

import (
	"fmt"

	"github.com/ahmadakhrom/go/goslash"
)

func main() {
	f, _ := goslash.Sum(5, 5, 5, 5, 5)
	fmt.Println(f)

	d, _ := goslash.Times(5,5,4)
	fmt.Println(d)
}
