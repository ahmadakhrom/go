package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNumber = errors.New("number must be more than zero")

func main() {

	fmt.Printf("type of : %T \n", ErrNumber) //-----*errors.errorString

	_, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(i int) (int, error) {
	if i < 0 {
		return 0, ErrNumber
	}
	return i * i, nil
}
