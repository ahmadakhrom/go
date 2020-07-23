package main

import (
	"fmt"
	"log"
)

type location struct {
	lat  string
	long string
	err  error
}

func (l location) Error() string {
	return fmt.Sprintf("error occurs %v %v %v", l.lat, l.long, l.err)
}

func sqrt(i int) (int, error) {

	if i < 0 {
		val := fmt.Errorf("error happened %v", i)
		return 0, location{"data lat", "data long", val}
	}
	return i * i, nil
}

func main() {
	v, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
