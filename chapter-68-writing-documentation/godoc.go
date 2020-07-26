package main

import (
	"chapter-68-writing-documentation/godark"
	"fmt"
)

func main() {
	f, _ := godark.Times(1, 2)
	fmt.Println(f)

	d, _ := godark.Sum(1, 3)
	fmt.Println(d)

	turn_years, _ := godark.Years(12)
	fmt.Println(turn_years, "dogs")

}