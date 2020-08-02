package main

import (
	"math"
)

func main() {
	//s := GetSum(10,11,19)
	//fmt.Println(s)

	//d :=GetAverage(10,35)
	//
	//fmt.Printf("%.1f\n",d)
}

//struct for method diameter and wide
type circle struct {
	size float64
}

//method diameter underlying from circle struct
func (c circle) diameter() float64 {
	return c.size / 2
}

//method wide underlying from circle struct
func (c circle) wide() float64 {
	return math.Phi * (math.Pow((c.size / 2), 2))
}

//get average from unlimited numbers are filled
func GetAverage(i ...float64) float64 {

	var a, b float64
	for _, v := range i {
		a += v
		b = a / 2
	}
	return b
}

//struct for func GetSum
type Answer struct {
	Guess  int
	result []int
}

//get summary from array int numbers are filled
func GetSum(i []int) int {

	f := 0
	for _, d := range i {
		f += d
	}
	return f
}

//get summary from unlimited int are filled
func GetTotal(i ...int) int {

	f := 0
	for _, d := range i {
		f += d
	}
	return f
}
