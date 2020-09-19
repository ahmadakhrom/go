package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/tpl.gohtml"))
}

func timeTo24h(t time.Time) string {
	return t.Format("15:04:05")
}

func dateToMDY(t time.Time) string  {
	return t.Format("Monday, part-02-Jan-2006")
}

var fm = template.FuncMap{
	"timeTo24h": timeTo24h,
	"dateToMDY" :dateToMDY,

}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalf("%err error occured", err)
	}
}
