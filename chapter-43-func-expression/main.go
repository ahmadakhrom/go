package main

import "fmt"

func main() {

	d := func() {
		fmt.Println("youre func in variable d")
	}

	d()

	f := func(numb int) {
		fmt.Println("numb is", numb)
	}

	f(33)
}
