package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml", "tpl2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) //printed gohtml.1
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil) //printed gohtml2
	if err != nil {
		log.Fatalln(err)
	}
}
