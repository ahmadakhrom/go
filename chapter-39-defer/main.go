package main

import "fmt"

func main() {

	defer fmt.Println("Loading....")

	x := Fullname("andy nugraha")
	y := Callme("andy")

	fmt.Println(x, y)
}

func Callme(s string) string {
	return fmt.Sprintln("call me is..", s)
}

func Fullname(s string) string {
	return fmt.Sprintln("my fullname is", s)
}
