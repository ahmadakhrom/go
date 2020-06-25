package main

import "fmt"

var a = "Hello world.."
var b = "Hello world"

func main() {
	fmt.Println(a)
	fmt.Printf("%T", a)

	d := []byte(a) //i try to convert to byte type
	fmt.Println(d)
	fmt.Printf("type of d is :  %T \n", d)

	for i, z := range d {
		fmt.Println(i, "character the desimal is", z)
	}
	for i, z := range a {
		fmt.Printf("index of a is : %d, and chararacter UTF 8 is : %#U \n", i, z)
	}
	for i, z := range a {
		fmt.Printf("index of a is : %d, and HEX is : %#x \n", i, z)

	}
}
