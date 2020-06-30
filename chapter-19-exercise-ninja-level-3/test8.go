package main

import "fmt"

func main() {

	switch {
	case !false:
		fmt.Println("false statement")

	case !true:
		fmt.Println("true statement")
	}
}
