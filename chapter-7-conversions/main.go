package main

import "fmt"

var a int

type blue int

var b blue

func main() {
	a = 100
	b = 45

	fmt.Println(a)
	fmt.Println(b)
	a = int(b)          //<== i can do this cause i was convert b TYPE of blue change to TYPE of int
	fmt.Printf("%T", b) //<== basicly b is stil TYPE of blue

}
