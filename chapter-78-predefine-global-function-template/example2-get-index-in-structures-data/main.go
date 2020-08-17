package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type Person struct {
	City []string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {
	city := []string{"Jakarta", "Bandung", "Bali", "Medan"}
	data := Person{city, 34}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("%err", err)
	}
}
