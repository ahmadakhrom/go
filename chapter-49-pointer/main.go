package main

import "fmt"

func main() {
	d := 12
	fmt.Println("memory addres of value d is ", &d)    // 0xc0000100b0
	fmt.Printf("and the type of pointer is %T \n", &d) //*int

	f := &d
	fmt.Println(*&f)

	*f = 32
	fmt.Println(*f)
	fmt.Println(&f)

}
