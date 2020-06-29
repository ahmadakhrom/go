package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	sedan_var := sedan{
		vehicle: vehicle{door: 4, color: "blue"},
		luxury:  true,
	}

	truck_var := truck{
		vehicle:   vehicle{door: 2, color: "red"},
		fourWheel: false,
	}

	fmt.Println(sedan_var)
	fmt.Println(truck_var)

	fmt.Println(sedan_var.color, sedan_var.door, sedan_var.luxury)
	fmt.Println(truck_var.color, truck_var.door, truck_var.fourWheel)

}
