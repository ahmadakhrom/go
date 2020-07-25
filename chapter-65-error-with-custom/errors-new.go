package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	val, err := sqrt(5)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result :", val)
}

func sqrt(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("number must be more than zero") //used custome error
	}
	return i * i, nil
}
