package main

import (
	"log"
	"os"
	"text/template"
)

type motogpRider struct {
	Name, Number, Manufacturer string
}

type winning struct {
	Podium int
	Motogp []motogpRider
}

type year struct {
	FirstYears, SecondYears, ThirdYears winning
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {

	data := year{
		FirstYears: winning{
			Podium: 1,
			Motogp: []motogpRider{
				{
					"Marq Marquez",
					"93",
					"Honda Repsol",
				},
			},
		},
		SecondYears: winning{
			Podium: 3,
			Motogp: []motogpRider{
				{
					"Andrea Dovizioso",
					"02",
					"Ducati MotoGP Racing",
				},
			},
		},
		ThirdYears: winning{
			Podium: 2,
			Motogp: []motogpRider{
				{
					"Valetino Rossi",
					"46",
					"Yamaha Fiat",
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalf("%err! error occured", err)
	}
}
