package main

import "fmt"

func main() {
	degree := 1.5
	if degree > 7.0 {
		fmt.Println("youre potentially indicated flu")
	} else if degree > 6.0 {
		fmt.Println("your temp body is normal")
	} else {
		fmt.Println("undefined")
	}
}
