package main

import "fmt"

func main() {
	a := []string{
		"james", "bond", "Shaken, not stirred",
	}

	b := []string{
		"Miss", "Moneypenny", "Helloooooo, James.",
	}

	c := [][]string{a, b}
	// fmt.Println(c)

	for i, xs := range c {
		fmt.Println("record", i)
		for i, value := range xs {
			fmt.Printf("\t index :%v value :%v \n", i, value)
		}
	}
}
