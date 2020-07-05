package main

import "fmt"

func main() {
	//anonymous func
	func(s string, a int) {
		fmt.Println("hallo", s, "umur kamu sekarang", a)
	}("andy nugraha"+",", 27)

}
