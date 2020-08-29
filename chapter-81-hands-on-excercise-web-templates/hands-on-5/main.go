package main

import (
	"html/template"
	"log"
	"os"
)

type items struct {
	Name, Desc string
	Price int
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
					NameOfMeals: "Breakfast",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 2",
							Desc: "desc..",
							Price: 15000,
						},
					},
				},
				Meals{
					NameOfMeals: "Launch",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
					},
				},
				Meals{
					NameOfMeals: "Dinner",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 2",
							Desc: "desc..",
							Price: 15000,
						},
					},
				},
			},
		},
		Restaurant{
			NameOfRestaurant: "European Food",
			Meals: []Meals{
				Meals{
					NameOfMeals: "Breakfast",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 2",
							Desc: "desc..",
							Price: 15000,
						},
					},
				},
				Meals{
					NameOfMeals: "Launch",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
					},
				},
				Meals{
					NameOfMeals: "Dinner",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 2",
							Desc: "desc..",
							Price: 15000,
						},
					},
				},
			},
		},
			Restaurant{
			NameOfRestaurant: "Asian Food",
			Meals: []Meals{
				Meals{
					NameOfMeals: "Breakfast",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 2",
							Desc: "desc..",
							Price: 15000,
						},
					},
				},
				Meals{
					NameOfMeals: "Launch",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
					},
				},
				Meals{
					NameOfMeals: "Dinner",
					Item: []items{
						items{
							Name: "dish 1",
							Desc: "desc..",
							Price: 14000,
						},
						items{
							Name: "dish 2",
							Desc: "desc..",
							Price: 15000,
						},
					},
				},
			},
		},
	}

		err := tpl.Execute(os.Stdout, data)
		if err != nil {
		log.Fatalf("%s", err)
	}
}
//restaurants -> []restaurant -> restaurant1 --> Name & []meals  -> item -> []items -> item1, item2, item3, etc,
//-----------------------------> restaurant2 --> Name & []meals  -> item -> []items -> item1, item2, item3, etc,
//-----------------------------> etc
