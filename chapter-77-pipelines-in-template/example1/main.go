package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template
var fm = template.FuncMap{
	"double":   double,
	"square":   square,
	"sqrtRoot": sqrtRoot,
}

func double(i int) int {
	return i + i
}

func square(i int) float64 {
	return math.Pow(float64(i), 2)
}

func sqrtRoot(i float64) float64 {
	return math.Sqrt(i)
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/tpl.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 5)
	if err != nil {
		log.Fatalf("%err error occured \n", err)
	}
}
