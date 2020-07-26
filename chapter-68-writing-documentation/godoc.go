package main

import (
	"chapter-68-writing-documentation/godark"
	"fmt"
//"g"
	)

func main() {
	f, _ :=godark.Sum(1, 2)
	fmt.Println(f)

	d, _ := godark.Times(1, 3)
	fmt.Println(d)

}