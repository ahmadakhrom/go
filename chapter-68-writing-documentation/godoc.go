package main

import (
	"fmt"

	"github.com/ahmadakhrom/go/goslash"
)

func main() {
	val, _ := goslash.Sum(1, 2, 3, 4, 5)

	fmt.Println(val)
}
