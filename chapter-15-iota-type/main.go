package main

import "fmt"

const (
	a = iota
	b
	c
	d
	e
	f
)

const (
	g = iota
	h
	i
)
const (
	_ = iota // setup increment numbers

	kb = 1 << (iota * 10) //sama dengan kb = 1 << (1 * 10)
	mb = 1 << (iota * 10) //sama dengan mb = 1 << (2 * 10)
	gb = 1 << (iota * 10)//sama dengan kb = 1 << (3 * 10)
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	fmt.Printf("on desimal : %d, on byte :\t\t %b\n", kb, kb)
	fmt.Printf("on desimal : %d, on byte :\t\t %b\n", mb, mb)
	fmt.Printf("on desimal : %d, on byte :\t %b\n", gb, gb)
}
