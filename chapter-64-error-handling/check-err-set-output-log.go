package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file1, err := os.Create("log.txt")
	defer file1.Close()
	log.SetOutput(file1)

	if err != nil {
		log.Fatal(err) // equivalent with os.Exit(1)
	}

	file2, err := os.Open("fileX.txt")
	defer file2.Close()

	if err != nil {
		log.Println("error", err)
	}

	fmt.Println("checking log.txt")

}
