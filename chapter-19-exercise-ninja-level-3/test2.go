package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println("No", i)
		for f := 1; f <= 1; f++ {
			fmt.Printf("\t%#U \n", i)
		}
	}
}
