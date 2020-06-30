package main

import "fmt"

func main() {

	//composite literal
	x := []int{1, 9, 12, 99, 3}
	fmt.Println(x)

	fmt.Println("\n")

	//index position values on var x
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	//knowing iteration of index position
	for i, x := range x {
		fmt.Println(i, "<== indexed on", x)
	}
}
