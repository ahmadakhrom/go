package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := []Person{
		{"James", 32},
		{"Moneypenny", 27},
		{"G", 54},
	}
	fmt.Println(p1)

	//marshaling
	resp, err := json.MarshalIndent(p1, "", "  ")
	if err != nil {
		fmt.Println("there was error")
		return
	}
	fmt.Println(string(resp))
}
