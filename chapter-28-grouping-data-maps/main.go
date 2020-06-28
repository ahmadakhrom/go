package main

import "fmt"

func main() {

	m := map[string]int{
		"indonesia": 17,
		"america":   4,
	}

	fmt.Println(m["indonesia"]) //oka
	fmt.Println(m["j"])         //not found coz no key on map "m"

	if res, err := m["indonesia"]; err {
		fmt.Println("it found, the value is :", res)
	}

	if res, err := m["germany"]; !err {
		fmt.Println("not found, the value is :", res)
	}

}
