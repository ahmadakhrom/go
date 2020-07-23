package main

import "fmt"

var a int

type blue float64

var b blue

func main() {
	a = 100
	b = 4.5

	fmt.Println(a)
	fmt.Printf("%T \n", a)

	fmt.Println("-----------------------------")

	fmt.Println(b)
	fmt.Printf("%T \n", b)

	fmt.Println("-----------------------------")

	b = blue(a) //conversion of data type
	fmt.Println(b)
	fmt.Printf("%T", b) //<== basicly b is stil TYPE of blue

}
