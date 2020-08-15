package main

import (
	"log"
	"os"
	"text/template"
)

	type M map[string]interface{}
	var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	data :=M{
		"name" :"Andy Nugraha",
		"age" : "27",
		"hobbies" : []string{"football","badminton","swiming"},
	}

	err := tpl.ExecuteTemplate(os.Stdout,"tpl2.gohtml",data)
	if err != nil{
		log.Fatal(err)
	}

}