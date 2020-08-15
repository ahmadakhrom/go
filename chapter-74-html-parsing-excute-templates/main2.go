package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	err := tpl.Execute(os.Stdout,nil)
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",nil)  //printed gohtml.1
	if err !=nil{
		log.Fatalln(err)
	}

		err = tpl.ExecuteTemplate(os.Stdout,"tpl2.gohtml",nil) //printed gohtml2
	if err !=nil{
		log.Fatalln(err)
	}
}