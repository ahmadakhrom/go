package main

import "fmt"

func main() {

	gotValue := func(f []int) int {
		if len(f) == 0 {
			return 0
		}
		if len(f) == 1 {
			return f[0]
		}
		return f[0] + f[len(f)-1] //1+7

	}

	rangesNumber := []int{1, 2, 3, 4, 5, 6, 7} //------ point of core
	d := calculate(gotValue, rangesNumber)     //how to get a value from first and last sequences from variabel X (func ""gotValue)
	fmt.Println(d)                             //and plus number from a func (calculate)
}

func calculate(f func(a []int) int, b []int) int {
	a := f(b)
	a++
	return a
}
