package main

import "fmt"

func main()  {
	s := "h"

	fmt.Println(s)
	fmt.Println("'h' convert to byte :",[]byte(s))
	y := s[0]
	fmt.Println("what the decimal of 'h' is :",y)
	fmt.Printf("what the type of 'h' : %T \n",y)
	fmt.Printf("what the binner of 'h' : %b \n",y)
	fmt.Printf("what the hexadesimal of 'h' : %#x \n",y)

	//you could check this out on
	//https://en.wikipedia.org/wiki/ASCII
	//https://www.belajarcpp.com/tutorial/sistem-bilangan/

}