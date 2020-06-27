package main

import "fmt"

const a = "x"

func main()  {
	b := a[0]
	fmt.Printf(" decimal 'x' is :\t %d \n binary 'x' is :\t %b \n hex 'x' is :\t\t %#x",b,b,b)
}