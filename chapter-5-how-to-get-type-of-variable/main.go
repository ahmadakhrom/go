package main

import "fmt"

//outside of body main function

var s string = "i'm have done mycourse before 2020"

func main() {

	fmt.Println(s)
	fmt.Printf("type of variable s is %T", s)

	//if i'm change my s value with int TYPE and then the result is crash.
	//coz you can't store var s to number oz its a int of TYPE
	s = 120
	fmt.Println(s)

}
