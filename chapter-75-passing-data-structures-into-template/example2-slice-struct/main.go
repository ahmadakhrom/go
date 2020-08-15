package main

import (
	"html/template"
	"log"
	"os"
)

type smartphone struct {
	Brand        string
	Feedback     string
	Manufacturer string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {

	iphone := smartphone{"iphone", "too expensive", "us"}
	samsung := smartphone{"samsung", "friendly", "south korea"}
	huawei := smartphone{"huawei", "expected", "taiwan"}

	data := []smartphone{iphone, samsung, huawei}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
