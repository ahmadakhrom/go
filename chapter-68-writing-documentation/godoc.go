package main

import (
	"fmt"

	"github.com/ahmadakhrom/go/gokill"
)

func main() {

	f, _ := gokill.Sum()(12, 2, 8, 100, 200)
	fmt.Println(f)

}
