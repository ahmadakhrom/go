package main

import "fmt"

var c int

func main()  {
	a := 10
	b := 44.8970
	c = 44.8970

	fmt.Printf("%T \n",a)
	fmt.Printf("%T \n",b)
	fmt.Printf("this doesn't run : %T \n",c)  //cause you before was decalre with int TYPE

}