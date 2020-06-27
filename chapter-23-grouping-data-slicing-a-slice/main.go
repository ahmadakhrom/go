package main

import "fmt"

func main() {

	//composite literal
	x := []int{1, 9, 12, 99, 3}

	//index position values on var x
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3:])  // we could got spesific slice and with specific position
	fmt.Println(x[4:5]) // we could got spesific slice and with specific position

	fmt.Println("\n----")

	//knowing iteration of index position, the way #1
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println("\n----")

	//knowing iteration of index position, the way #2
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
