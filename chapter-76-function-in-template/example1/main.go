package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

type smartphone struct {
	Brand        string
	Feedback     string
	Manufacturer string
}

type owner struct {
	Name       string
	Profession string
}

var fm = template.FuncMap{
	"lc": strings.ToLower,
	"us": upperString,
}

func init() {
	//define existing template with "template.new" to using funcmap.
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/tpl.gohtml"))
}

func upperString(s string) string {
	s = strings.ToUpper(s)
	return s
}

func main() {

	iphone := smartphone{"IPHONE", "TOO EXPENSIVE", "US"}
	samsung := smartphone{"samsung", "friendly", "south korea"}
	huawei := smartphone{"huawei", "expected", "taiwan"}

	people1 := owner{"andy", "teacher"}
	people2 := owner{"angga", "bmx bike racer"}
	people3 := owner{"budi", "mechanic"}

	smartphones := []smartphone{iphone, samsung, huawei}
	owners := []owner{people1, people2, people3}

	data := struct {
		Device  []smartphone
		Details []owner
	}{
		smartphones,
		owners,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
