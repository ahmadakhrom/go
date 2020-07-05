package main

import "fmt"

func main() {

	d := []int{2, 4, 5, 7, 4, 3, 3, 54, 5, 4} //make a slice
	xx := car(d...)                           //"..." make it free how many valus your stored
	fmt.Println("Count of your car is", xx)   //its a diferent technic but same result on previeous chapter

}

func car(list ...int) int {
	fmt.Println(list)
	fmt.Printf("%T \n", list)
	fmt.Println("length", len(list))
	fmt.Println("capacity", cap(list))

	numb := 0
	for i, v := range list {
		numb += v
		fmt.Println(i, "andy have a car as much : ", v)
	}

	fmt.Println("Count of your car is", numb)
	return numb
}
