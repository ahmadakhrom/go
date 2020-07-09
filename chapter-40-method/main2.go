package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
	length int
}

type shape interface {  //made for stored any method where you created
	area() float64
	areb() int
}

func (c *circle) area() float64 {		//make method #1
	return math.Pi * c.radius * c.radius
}

func (c *circle) areb() int { ////make method #2
	return  100 * c.length
}

//func (c *circle) areb() int { ////make method #your next method..
//	return  100 * c.length
//}


func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5,5}
	// info(c)
	fmt.Println(c.area())
	fmt.Println(c.areb())
}
