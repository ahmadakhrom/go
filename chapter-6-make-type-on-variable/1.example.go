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
	//a=b   <== you cannot do this coz : b (type blue) as type int in assignment
	fmt.Printf("%T", b)

}
