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

type owner struct {
	Name       string
	Profession string
	}

type items struct {
	Device []smartphone
	Work   []owner
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {

	iphone := smartphone{"iphone", "too expensive", "us"}
	samsung := smartphone{"samsung", "friendly", "south korea"}
	huawei := smartphone{"huawei", "expected", "taiwan"}

	people1 := owner{"andy", "teacher"}
	people2 := owner{"angga", "bmx bike racer"}
	people3 := owner{"budi", "mechanic"}

	smartphones := []smartphone{iphone, samsung, huawei}
	owners := []owner{people1, people2, people3}

	data := items{smartphones, owners}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
