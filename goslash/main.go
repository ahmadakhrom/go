package main

import (
	"fmt"

	"github.com/ahmadakhrom/go/goslash")

func main() {
	d, _ := goslash.Sum(1,2)
	fmt.Println(d)
}