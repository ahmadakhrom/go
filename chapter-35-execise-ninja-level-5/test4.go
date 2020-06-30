package main

import "fmt"

func main() {
	a := struct {
		gender  int
		name    string
		vehicle map[string]int // itu will show like : "toyota" : 1
		son     []string       //it will format for insert array son
	}{
		gender:  1,
		name:    "andy nugraha",
		vehicle: map[string]int{"honda": 1, "volkswagen": 2},
		son:     []string{"garry", "david", "barry"},
	}
	fmt.Println(a.name)
	fmt.Println(a.gender)

	for i, v := range a.vehicle {
		fmt.Println("he had a car is", i, "as much", v) //ini "i" nya bukan indeks melainkan KEY yang ada pada map vehicle
	}

	for i, val := range a.son { //"i" nya adalah indeks,
		fmt.Println(" has son : ", i, val)

	}

}
