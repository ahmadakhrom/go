package main

import "fmt"

func main() {

	//make(TYPE , length , capacity) //type variable must be Int
	a := make([]int,10,125)

	fmt.Println(a) // index position allocated to store an array
	fmt.Println(len(a)) //length of array
	fmt.Println(cap(a)) //capacity of array

	a[3] = 50
	a[7] = 70
	
	fmt.Println(a)
	fmt.Println(append(a,88))

	
}