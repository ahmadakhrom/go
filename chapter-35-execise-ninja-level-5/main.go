package main

import "fmt"

func main() {

	type name string          //we make a certainly TYPE
	var son name = "john doe" //and we could assign a value to variable has a TYPE that we was created before

	const other string = "Gerrald" // these constant variable

	fmt.Println(son)
	fmt.Println(other)
}
