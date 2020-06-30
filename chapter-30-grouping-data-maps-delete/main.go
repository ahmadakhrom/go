package main

import "fmt"

func main() {

	m := map[string]int{
		"indonesia": 17,
		"america":   4,
	}

	// could print all
	fmt.Println(m)

	// I try to declare but not listed on variable "m"
	fmt.Println(m["germany"]) //null result

	//I try to remove key on variable "m"
	delete(m, "america")
	delete(m, "germany")

	//try to print varible "m"
	fmt.Println(m) //this will print just 1 key is : indonesia

	//try to delete all key on "m"
	if _, err := m["indonesia"]; err {
		fmt.Println("run now", m)
		delete(m, "indonesia")
	}
	fmt.Println("run now", m)
}
