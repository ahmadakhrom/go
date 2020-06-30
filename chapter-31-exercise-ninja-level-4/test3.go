package main

import "fmt"

func main() {

	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(a[:5])  //[42 43 44 45 46]
	fmt.Println(a[5:])  //[47 48 49 50 51]
	fmt.Println(a[2:7]) //[44 45 46 47 48]
	fmt.Println(a[1:6]) //[43 44 45 46 47]

}
