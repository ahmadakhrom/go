package main

import (
	"text/template"
	"log"
	"os"
)

type page struct {
	Title string
	Heading string
	Message	string
}

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main()  {
	data := page{
		"How to learn Golang",
		"Using Template Package",
		`<script>alert("hello world");</script>`,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("%s error occured",err)
	}
}