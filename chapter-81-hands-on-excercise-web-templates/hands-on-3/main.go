package main

import (
	"html/template"
	"log"
	"os"
)

type restaurant struct {
	Name string
	Breakfast []string
	Launchs []string
	Dinner  []string
}

type californiaHotel struct {
	Name, Address, City, Zip string
	Restaurants []restaurant
}

type region struct {
	CaliforniaHotels []californiaHotel
	RegionName string
}

type regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	data := regions{
		region{
			RegionName: "Southern",
			CaliforniaHotels: []californiaHotel{
				{
					Name: "Hyatt Regency Monterey Hotel",
					Address: "1 Old Golf Course Rd, Monterey, CA , United States",
					City :"California",
					Zip:"93940",
					Restaurants:[]restaurant{
						{
							"Sangrila Resto",
							[]string{"s", "s", "s"},
							[]string{"s", "s", "s"},
							[]string{"s", "s", "s"},
						},
						{
							 "Sangrila2 Resto",
							 []string{"s", "s", "s"},
							[]string{"s", "s", "s"},
							[]string{"s", "s", "s"},
						},
					},
				},
				{
					Name: "Hyatt Regency Monterey Hotel",
					Address: "1 Old Golf Course Rd, Monterey, CA , United States",
					City: "California",
					Zip: "93940",
					Restaurants:[]restaurant{
						{
							"California Lake",
							[]string{"f","f"},
							[]string{"f","f"},
							[]string{"f","f"},
						},
						{
							"California Lake2",
							[]string{"g","g"},
							[]string{"g","g"},
							[]string{"g","g"},
						},
					},
				},
			},
		},
		region{
			RegionName: "Central",
			CaliforniaHotels: []californiaHotel{
				{
					Name: "Black Hawk Lodge Hotel",
					Address: "5729+54 Yosemite Lakes, California, United States",
					City: "California",
					Zip: "93645",
					Restaurants:[]restaurant{
						{
							"Golden Swiss Resto",
							[]string{"d","d"},
							[]string{"d","d"},
							[]string{"d","d"},
						},
					},
				},

				{
					Name: "Black Hawk Lodge Hotel",
					Address: "5729+54 Yosemite Lakes, California, United States",
					City:"California",
					Zip: "93645",
					Restaurants:[]restaurant{
						{
							"Golden Swiss Resto2",
							[]string{"s","s"},
							[]string{"s","s"},
							[]string{"s","s"},
						},
					},
				},
			},
		},
		region{
			RegionName: "Northern",
			CaliforniaHotels: []californiaHotel{
				{
					Name: "Bella Vista Bed & Breakfast Hotel",
					Address: "581 Cold Springs Rd, Placerville, CA, United States",
					City:"California",
					Zip: "95667",
					Restaurants:[]restaurant{
						{
							"Asian Food",
							[]string{"s","s"},
							[]string{"s","s"},
							[]string{"s","s"},
						},
					},
				},
				{
					Name: "Bella Vista Bed & Breakfast Hotel",
					Address: "581 Cold Springs Rd, Placerville, CA, United States",
					City: "California",
					Zip: "95667",
					Restaurants:[]restaurant{
						{
							"Asian Food2",
							[]string{"a","a"},
							[]string{"a","a"},
							[]string{"a","a"},
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("%err error occured", err)
	}

}
