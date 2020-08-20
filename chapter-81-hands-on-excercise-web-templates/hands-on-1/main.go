package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term   string
	Course []course
}

type year struct {
	AcaYear              string
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//single of year
	//data := year{
	//	AcaYear: "2019",
	//	Fall:    semester{
	//		Term: "Fall",
	//		Course: []course{
	//			{
	//				" CSCI-40", "Introduction to Programming in Go", "4",
	//			},
	//			{
	//				"CSCI-130", "Introduction to Web Programming with Go", "4",
	//			},
	//			{
	//				"CSCI-140", "Mobile Apps Using Go", "4",
	//			},
	//		},
	//	},
	//Spring:    semester{},
	//Summer:    semester{},
	//}

	//multiple of year range
	data := []year{
		year{
			AcaYear: "2018",
			Fall: semester{
				Term:   "Fall Season",
				Course: []course{course{"TEST1", "TEST1", "TEST1"}},
			},
			Spring: semester{
				Term:   "Spring Season",
				Course: []course{course{"TEST2", "TEST2", "TEST2"}},
			},
			Summer: semester{
				Term:   "Summer Season",
				Course: []course{course{"TEST3", "TEST3", "TEST3"}},
			},
		},
		year{
			AcaYear: "2019",
			Fall: semester{
				Term:   "Fall Season",
				Course: []course{course{"TEST4", "TEST4", "TEST4"}},
			},
			Spring: semester{
				Term:   "Spring Season",
				Course: []course{course{"TEST5", "TEST5", "TEST5"}},
			},
			Summer: semester{
				Term:   "Summer Season",
				Course: []course{course{"TEST6", "TEST6", "TEST6"}},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalf("%err error occured", err)
	}
}
