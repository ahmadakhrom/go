//Method sets determine what methods attach to a TYPE.
//It is exactly what the name says: What is the set of methods for a given type? That is its method set.

package main

import "fmt"

func main() {
	d := 12
	fmt.Println("memory addres od value d is ", &d)    // 0xc0000100b0
	fmt.Printf("and the type of pointer is %T \n", &d) //*unt

	f := &d
	fmt.Println(*&f)

	*f = 32
	fmt.Println(*f)
	fmt.Println(&f)

}
