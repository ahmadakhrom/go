package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["John doe"] = []string{`one`, `two`, `three`}

	for i, v := range m {
		fmt.Println("name is", i)
		for in, val := range v {
			fmt.Println("\tfav is", in, val)
		}
	}

}
