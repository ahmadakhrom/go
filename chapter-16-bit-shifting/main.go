package main

import "fmt"

func main()  {
	const (
		a = iota

		b
		c = b * a
	)
fmt.Println("",a)
fmt.Println("",b)
fmt.Println("",c)
}