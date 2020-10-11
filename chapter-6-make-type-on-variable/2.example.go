package main

import "fmt"

func main() {
	foo := make([]string, 0,10)
	f := append(foo,"bar","john")
	fmt.Println(f[0:2])
}
