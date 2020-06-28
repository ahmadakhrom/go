package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["john_doe"] = []string{`one`, `two`, `three`}

	delete(m, `bond_james`)

	for i, v := range m {

		fmt.Println("name is", i)
		for in, val := range v {
			fmt.Println("\t fav is", in, val)
		}
	}
	fmt.Println("bond_james map deleted")
}
