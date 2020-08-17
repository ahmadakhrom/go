package main

import (
	"html/template"
	"log"
	"os"
)

type lastNightUcl struct {
	Lyon    int
	Citizen int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {

	score := lastNightUcl{
		3,
		1,
	}

	err := tpl.Execute(os.Stdout, score)
	if err != nil {
		log.Fatalf("%err", err)
	}
}
