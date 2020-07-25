package main

import (
	"fmt"

	"github.com/ahmadakhrom/goflash"
)

func main() {
	f, _ := goflash.Sum(1,2)
	fmt.Println(f)

	d, _ := goflash.Times(1,3)
	fmt.Println(d)
}
