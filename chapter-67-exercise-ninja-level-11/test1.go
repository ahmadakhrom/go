package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := Person{
		"James",
		"Bond",
		[]string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	byteResp, err := json.Marshal(p1)
	if err == nil {
		log.Fatalln("error marshal", err)
	}

	fmt.Println(string(byteResp))
}
