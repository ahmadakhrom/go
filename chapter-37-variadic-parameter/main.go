package main

import "fmt"

func main() {

	d := car(2, 4, 5, 7, 4, 3, 3, 54, 5, 4)
	fmt.Println("Count of your car is", d)

}

func car(list ...int) int {
	fmt.Println(list)
	fmt.Printf("%T \n", list)

	numb := 0
	for i, v := range list {
		numb += v
		fmt.Println(i, "andy have a car as much : ", v)
	}

	fmt.Println("Count of your car is", numb)
	return numb
}
