package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {
	xs := []string{"Jakarta", "Bandung", "Bali", "Medan"}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalf("%err", err)
	}
}
