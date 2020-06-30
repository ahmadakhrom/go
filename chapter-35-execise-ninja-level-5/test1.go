package main

import "fmt"

//#1
type person struct {
	first_name        string
	last_name         string
	ice_cream_flavors []string
}

func main() {

	a := person{
		first_name: "andy",
		last_name:  "nugraha",
		ice_cream_flavors: []string{
			"oreo vanilla",
			"sweet choco",
		},
	}

	fmt.Println(a.first_name, a.last_name, a.ice_cream_flavors)
	for i, v := range a.ice_cream_flavors {
		fmt.Println("he likes a", i, v)
	}
 




}
