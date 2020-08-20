package main

import (
	"html/template"
	"log"
	"os"
)

type items struct {
	Name, Desc, Price string
}

type Meals struct {
	NameOfMeals string
	Item []items
}

type  Restaurant struct {
	NameOfRestaurant string
	Meals []Meals
}

type restaurants []Restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	data := restaurants{
		Restaurant{
			NameOfRestaurant: "Chinese Food",
			Meals: []Meals{
				Meals{
					NameOfMeals: "",
					Item: []items{
						items{
							Name: "",
							Desc: "",
							Price: "",
						},
						items{
							Name: "",
							Desc: "",
							Price: "",
						},
					},
				},
				Meals{
					NameOfMeals: "",
					Item: []items{
						items{
							Name: "",
							Desc: "",
							Price: "",
						},
						items{
							Name: "",
							Desc: "",
							Price: "",
						},
					},
				},
			},
		},
		Restaurant{
			NameOfRestaurant: "European Food",
			Meals: []Meals{
				Meals{
					NameOfMeals: "",
					Item: []items{
						items{
							Name: "",
							Desc: "",
							Price: "",
						},
						items{
							Name: "",
							Desc: "",
							Price: "",
						},
					},
				},
				Meals{
					NameOfMeals: "",
					Item: []items{
						items{
							Name: "",
							Desc: "",
							Price: "",
						},
						items{
							Name: "",
							Desc: "",
							Price: "",
						},
					},
				},
			},
		},
	}

		err := tpl.Execute(os.Stdout, data)
		if err != nil {
		log.Fatalf("%err", err)
	}
}
//restaurants -> []restaurant -> restaurant1 --> Name & []meals  -> item -> []items -> item1, item2, item3, etc,
//-----------------------------> restaurant2 --> Name & []meals  -> item -> []items -> item1, item2, item3, etc,
//-----------------------------> etc
