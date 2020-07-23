package main

import (
	"fmt"
)

//var ErrorData = errors.New("errorrrrr..")

type customerErr struct {
	Name string
}

func (ce customerErr) Error() string {
	return fmt.Sprintf("error happenend : %v", ce.Name)
}

func main() {

	c1 := customerErr{
		"alexis",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println(e, e.(customerErr).Name)
	fmt.Println(e, e.(customerErr).Name) //asertion technique
}
