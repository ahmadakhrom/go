package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())

}

func main() {

	square := square{
		length: 15,
	}

	circle := circle{
		radius: 10.55,
	}

	info(circle) //result of formula circle
	info(square) //result of formula squre

}
