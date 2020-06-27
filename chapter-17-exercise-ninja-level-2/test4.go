package main

import "fmt"

func main() {
	b := 9
	fmt.Printf("decimal of b : %d \nbinary of b : %b \nhex of b : %#x\n\n", b, b, b) //print that variable b

	c := b << 1                                                                  //shifts the bits of that int over 1 position to the left
	fmt.Printf("decimal of c : %d \nbinary of c : %b \nhex of c : %#x", c, c, c) //print that variable c

}
