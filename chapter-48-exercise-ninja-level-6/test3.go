package main

import "fmt"

func main() {

	defer foo() //hello foo
	bar()       //hello bar

}

func foo() {
	fmt.Println("hello foo")
}

func bar() {
	fmt.Println("hello bar")
}
