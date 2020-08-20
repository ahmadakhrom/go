package main

import (
	"html/template"
	"log"
	"os"
)

type item struct {
	Name, Desc, Price string
}

type Meals struct {
	Meal string //Breakfast, Launch and Dinner
	Menu []item
}

type Restaurant []Meals


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	data :=Restaurant{
		Meals{
			Meal: "Breaksfast",
			Menu :[]item{
				{"2", "2", "2"},
				{"2", "2", "2"},
			},
		},
		Meals{
			Meal: "Launch",
			Menu :[]item{
				{"3", "3", "3"},
				{"3", "3", "3"},
			},
		},
		Meals{
			Meal: "Dinner",
			Menu :[]item{
				{"4", "4", "4"},
				{"4", "4", "4"},
			},
		},
	}



			//Restaurant{
			//	Meal: "Launch",
			//	Menu: []item{
			//		{"Chia Pudding", "Bread Upma", "Bread Upma"},
			//		{"Scallion Pancakes", "Banana Pancake", "Bread Upma"},
			//		{"Aloo Paratha", "Paratha varieties", "Eggless Pancakes"},
			//	},
			//},
			//Restaurant{
			//	Meal: "Dinner",
			//	Menu: []item{
			//		{"Homemade Pizza", "Twice-Baked Potatoes", "Chicken Pot Pie"},
			//		{"Homemade Corn Dogs", "Potato-Topped Casserole", "Calzones"},
			//		{"French Bread Pizza", "Enchiladas", "Enchiladas2"},
			//	},

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("%err", err)
	}
}

//restaurant --> Breakfast --> []Menu -->Desc
//									  --> Price

// 			--> Launch --> []Menu -->Desc
//									  --> Price

// 			--> Launch --> []Menu -->Desc
//									  --> Price
