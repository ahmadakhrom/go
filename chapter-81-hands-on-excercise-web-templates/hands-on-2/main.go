package main

import (
	"html/template"
	"log"
	"os"
)

type californiaHotel struct {
	Name, Address, City, Zip string
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
					"Hyatt Regency Monterey Hotel",
					"1 Old Golf Course Rd, Monterey, CA , United States",
					"California",
					"93940",
				},
				{
					"Hyatt Regency Monterey Hotel",
					"1 Old Golf Course Rd, Monterey, CA , United States",
					"California",
					"93940",
				},
			},
		},
		region{
			RegionName: "Central",
			CaliforniaHotels: []californiaHotel{
				{
					"Black Hawk Lodge Hotel",
					"5729+54 Yosemite Lakes, California, United States",
					"California",
					"93645",
				},
				{
					"Black Hawk Lodge Hotel",
					"5729+54 Yosemite Lakes, California, United States",
					"California",
					"93645",
				},
			},
		},
		region{
			RegionName: "Northern",
			CaliforniaHotels: []californiaHotel{
				{
					"Bella Vista Bed & Breakfast Hotel",
					"581 Cold Springs Rd, Placerville, CA, United States",
					"California",
					"95667",
				},
				{
					"Bella Vista Bed & Breakfast Hotel",
					"581 Cold Springs Rd, Placerville, CA, United States",
					"California",
					"95667",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("%err error occured", err)
	}

}
