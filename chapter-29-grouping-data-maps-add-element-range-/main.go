package main

import "fmt"

func main() {

	m := map[string]int{
		"indonesia": 17,
		"america":   4,
	}

	m["germany"] = 19 // I tried to add a value, coz I was declare before to variable "m"

	for i, v := range m {
		fmt.Println(i, v)
	}

}
