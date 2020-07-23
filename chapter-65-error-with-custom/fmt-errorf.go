package main

import (
	"fmt"
	"log"
)

func main() {
	val, err := sqrt(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("cannot use negative numbers %v", f)
	}

	return f * f, nil
}
