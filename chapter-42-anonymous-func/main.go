package main

import "fmt"

func main() {

	//func foo
	foo()

	//anonymous func
	func() {
		fmt.Println("this anonymous func with empty identifier")

	}()

	//anonymous func with identifier
	func(s string) {
		fmt.Println("this anonymous func with identifier string and value is :", s)
	}("andy nugraha")
}

//func defined
func foo() {
	fmt.Println("this a defined func")
}
