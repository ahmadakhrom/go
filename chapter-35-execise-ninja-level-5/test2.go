package main

import "fmt"

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

	fmt.Println(a.first_name, a.last_name)
	for i, v := range a.ice_cream_flavors {
		fmt.Println("\t", i, v)
	}

	x := map[string]person{
		a.last_name: a,
	}
	fmt.Println(x)

	for _, v := range x {
		fmt.Println("Nama depan:", v.first_name, "Nama Belakang:", v.last_name)
		for _, val := range v.ice_cream_flavors {
			fmt.Println("flavors", val)
		}
	}

}
