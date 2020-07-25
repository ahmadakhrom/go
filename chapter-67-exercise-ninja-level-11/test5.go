package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("error happenend %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		errs := fmt.Errorf("Error Happenend ")
		return 0, sqrtError{"0", "0", errs}
	}
	return f * f, nil
}

func main() {
	v, err := sqrt(10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v, err)
}
