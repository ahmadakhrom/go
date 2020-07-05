package main

import "fmt"

func main() {
	resp, err := foo()
	fmt.Println(resp)
	fmt.Println(err)

}

func foo() (string, bool) {
	s := "hello world, i was retun this string to func foo"
	f := true
	return s, f
}
