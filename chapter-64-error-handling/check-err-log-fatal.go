package main

import (
	"log"
	"os"
)

func main() {

	_, err := os.Open("newfile1.txt")
	if err != nil {
		log.Fatal(err) // equivalent with os.Exit(1)
	}
}
