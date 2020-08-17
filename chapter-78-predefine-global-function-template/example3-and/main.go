package main

import (
	"log"
	"os"
	"text/template"
)

type fundraiser struct {
	Founder      string
	Organization string
	Status       bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {
	first := fundraiser{
		"Andy F. Noya",
		"BenihBaik.com",
		true,
	}

	second := fundraiser{
		"M. Alfatih Timur",
		"KitaBisa.com",
		false,
	}

	third := fundraiser{
		"Parni Hadi, Haidar Bagir, Sinansari Ecip, dan Erie Sudewo",
		"DompetDhuafa.org",
		true,
	}

	fourth := fundraiser{
		"Peter Setiawan Shearer",
		"Rantang Hati",
		false,
	}

	data := []fundraiser{first, second, third, fourth}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("%err", err)
	}
}
