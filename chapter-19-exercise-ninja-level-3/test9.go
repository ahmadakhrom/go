package main

import "fmt"

func main() {
	a := "favSport"

	switch a {
	case "favSport":
		fmt.Println("yes..")
		//fallthrough //skip matching and print force to print result
	case "favSport2":
		fmt.Println("no it's..")
		//fallthrough
	case "favSport3":
		fmt.Println("no too..")

	default:
		fmt.Println("nothing..")
	}
}
