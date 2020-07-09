package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {

	dataPerson := []Person{
		{
			"James",
			"Bond",
			32,
			[]string{
				"James, it is soo good to see you",
				"Would you like me to take care of that for you, James?",
				"I would really prefer to be a secret agent myself."},
		},
		{
			"Miss",
			"MennyPenny",
			27,
			[]string{
				"James, it is soo good to see you",
				"Would you like me to take care of that for you, James?",
				"I would really prefer to be a secret agent myself."},
		},
		{
			"Andi",
			"Nugraha",
			27,
			[]string{
				"James, it is soo good to see you",
				"Would you like me to take care of that for you, James?",
				"I would really prefer to be a secret agent myself."},
		},
	}
	// fmt.Println(dataPerson[0])
	// fmt.Println(dataPerson[1])
	// fmt.Println(dataPerson[2])

	//	encoded := json.NewEncoder(os.Stdout).Encode(dataPerson)
	encoded := json.NewEncoder(os.Stdout)
	encoded.Encode(dataPerson)

	fmt.Fprintln(os.Stdout, encoded)

}
