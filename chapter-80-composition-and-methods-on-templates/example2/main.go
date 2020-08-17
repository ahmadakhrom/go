package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type person struct {
	Artist     string
	BeautyRate int
}

func (p person) Bad() int {
	return 65
}

func (p person) Standard() int {
	return p.BeautyRate * 7
}

func (p person) Amazing(x int) int {
	return x * 9
}

func (p person) CalcRate(x int) int {
	return x * 10
}

func (p person) Loop(x int) int {
	var i int
	for i=1; i<x; i++{
		fmt.Print(i)
	}
	return i
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {
	data := person{
		"Agnes Monica",
		8,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("%err error occured", err)
	}
}
