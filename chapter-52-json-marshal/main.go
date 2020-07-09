package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int
	Name string
	Age  int
}

func main() {

	champion1 := student{
		ID:   1,
		Name: "Andy Nugraha",
		Age:  27,
	}

	champion2 := student{
		ID:   2,
		Name: "Budi Budiman",
		Age:  26,
	}

	allChampions := []student{
		champion1,
		champion2,
	}

	//func Marshal(v interface{}) ([]byte, error)	
	resp, err := json.MarshalIndent(allChampions, " ", "  ")
	if err != nil {
		fmt.Println("error marshaling..")
	}

	fmt.Println(champion1)
	fmt.Println(champion2)
	fmt.Println(string(resp))
}
